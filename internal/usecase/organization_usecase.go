package usecase

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xerror"
	"inbody-ocr-backend/internal/usecase/response"
)

type organizationUsecase struct {
	repo          repository.OrganizationRepository
	userRepo      repository.UserRepository
	imageDataRepo repository.ImageDataRepository
	tokenService  service.TokenService
	ulidService   service.ULIDService
}

func NewOrganizationUsecase(repo repository.OrganizationRepository, userRepo repository.UserRepository, imageDataRepo repository.ImageDataRepository, tokenService service.TokenService, ulidService service.ULIDService) OrganizationUsecase {
	return &organizationUsecase{
		repo:          repo,
		userRepo:      userRepo,
		imageDataRepo: imageDataRepo,
		tokenService:  tokenService,
		ulidService:   ulidService,
	}
}

// signupも同時に行う
func (uc *organizationUsecase) CreateOrganization(userName, email, password, orgName string) (*response.CreateOrganizationResponse, error) {
	exists, err := uc.userRepo.UserExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, xerror.ErrEmailAlreadyExists
	}

	org := entity.Organization{
		ID:   uc.ulidService.GenerateULID(),
		Name: orgName,
	}

	hashedPassword, err := uc.userRepo.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ID:             uc.ulidService.GenerateULID(),
		Name:           userName,
		Email:          email,
		Password:       hashedPassword,
		OrganizationID: org.ID,
		Role:           entity.OrganizationRoleOwner,
	}

	organization, err := uc.repo.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	user, err = uc.userRepo.CreateUser(*user)
	if err != nil {
		// ユーザのいない組織が作成されるのを防ぐためにロールバック
		rollbackErr := uc.userRepo.DeleteUser(user.ID)
		if rollbackErr != nil {
			return nil, rollbackErr
		}

		return nil, err
	}

	token, err := uc.tokenService.GenerateTokenFromID(user.ID, user.OrganizationID)
	if err != nil {
		return nil, err
	}
	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewCreateOrganizationResponse(*organization, token, user.ID, user.Name, *exp)
}

func (uc *organizationUsecase) GetAllMembers(orgID string) (*response.GetAllMembersResponse, error) {
	users, err := uc.repo.GetMember(orgID)
	if err != nil {
		return nil, err
	}

	return response.NewGetAllMembersResponse(users)
}

func (uc *organizationUsecase) UpdateRole(updateUserID string, role entity.OrganizationRole, requestUser *entity.User) (*response.UpdateRoleResponse, error) {
	// ownerから別のロールへの変更は不可
	updateUser, err := uc.userRepo.FindByID(updateUserID)
	if err != nil {
		return nil, err
	}
	if updateUser.Role == "owner" {
		return nil, xerror.ErrCannotUpdateOwnerRole
	}

	err = uc.userRepo.UpdateRole(updateUserID, role)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uc.userRepo.FindByID(updateUserID)
	if err != nil {
		return nil, err
	}

	return response.NewUpdateRoleResponse(*updatedUser)
}

func (uc *organizationUsecase) DeleteMember(deleteUserID string, requestUser *entity.User) (*response.DeleteMemberResponse, error) {
	// ownerは削除不可
	deleteUser, err := uc.userRepo.FindByID(deleteUserID)
	if err != nil {
		return nil, err
	}

	if deleteUser.Role == "owner" {
		return nil, xerror.ErrCannotDeleteOwner
	}

	err = uc.userRepo.DeleteUser(deleteUserID)
	if err != nil {
		return nil, err
	}

	return response.NewDeleteMemberResponse()
}
