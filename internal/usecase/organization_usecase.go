package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
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
		logger.Error("CreateOrganization", "func", "UserExists()", "error", err.Error())
		return nil, err
	}
	if exists {
		logger.Error("CreateOrganization", "func", "UserExists()", "error", "email already exists")
		return nil, fmt.Errorf("email already exists")
	}

	org := entity.Organization{
		ID:   uc.ulidService.GenerateULID(),
		Name: orgName,
	}

	hashedPassword, err := uc.userRepo.HashPassword(password)
	if err != nil {
		logger.Error("CreateOrganization", "func", "HashPassword()", "error", err.Error())
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
		logger.Error("CreateOrganization", "func", "CreateOrganization()", "error", err.Error())
		return nil, err
	}

	user, err = uc.userRepo.CreateUser(*user)
	if err != nil {
		// ユーザのいない組織が作成されるのを防ぐためにロールバック
		rollbackErr := uc.userRepo.DeleteUser(user.ID)
		if rollbackErr != nil {
			logger.Error("CreateOrganization", "func", "DeleteUser()", "error", rollbackErr.Error())
			return nil, rollbackErr
		}

		logger.Error("CreateOrganization", "func", "CreateUser()", "error", err.Error())
		return nil, err
	}

	token, err := uc.tokenService.GenerateTokenFromID(user.ID, user.OrganizationID)
	if err != nil {
		logger.Error("CreateOrganization", "func", "GenerateTokenFromID", "error", err.Error())
		return nil, err
	}
	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		logger.Error("CreateOrganization", "func", "ExtractExpFromToken", "error", err.Error())
		return nil, err
	}

	return response.NewCreateOrganizationResponse(*organization, token, user.ID, user.Name, *exp)
}

func (uc *organizationUsecase) GetAllMembers(orgID string) (*response.GetAllMembersResponse, error) {
	users, err := uc.repo.GetMember(orgID)
	if err != nil {
		logger.Error("GetAllMembers", "func", "GetMember()", "error", err.Error())
		return nil, err
	}

	return response.NewGetAllMembersResponse(users)
}

func (uc *organizationUsecase) UpdateRole(updateUserID string, role entity.OrganizationRole, orgID, requestUserID string) (*response.UpdateRoleResponse, error) {
	// memberは編集権限を持たない
	requestUser, err := uc.userRepo.FindByID(requestUserID)
	if err != nil {
		logger.Error("UpdateRole", "func", "FindByID()", "error", err.Error())
		return nil, err
	}
	if requestUser.Role == "member" {
		logger.Error("UpdateRole", "error", "user is not admin")
		return nil, fmt.Errorf("user is not admin")
	}

	// ownerから別のロールへの変更は不可
	updateUser, err := uc.userRepo.FindByID(updateUserID)
	if err != nil {
		logger.Error("UpdateRole", "func", "FindByID()", "error", err.Error())
		return nil, err
	}
	if updateUser.Role == "owner" {
		logger.Error("UpdateRole", "error", "cannot update owner role")
		return nil, fmt.Errorf("cannot update owner role")
	}
	
	err = uc.userRepo.UpdateRole(updateUserID, role)
	if err != nil {
		logger.Error("UpdateRole", "func", "UpdateRole()", "error", err.Error())
		return nil, err
	}

	updatedUser, err := uc.userRepo.FindByID(updateUserID)
	if err != nil {
		logger.Error("UpdateRole", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	return response.NewUpdateRoleResponse(*updatedUser)
}

func (uc *organizationUsecase) DeleteMember(deleteUserID, orgID, requestUserID string) (*response.DeleteMemberResponse, error) {
	// memberは削除権限を持たない
	requestUser, err := uc.userRepo.FindByID(requestUserID)
	if err != nil {
		logger.Error("DeleteMember", "func", "FindByID()", "error", err.Error())
		return nil, err
	}
	if requestUser.Role == "member" {
		logger.Error("DeleteMember", "error", "user is not admin")
		return nil, fmt.Errorf("user is not admin")
	}

	// ownerは削除不可
	deleteUser, err := uc.userRepo.FindByID(deleteUserID)
	if err != nil {
		logger.Error("DeleteMember", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	if deleteUser.Role == "owner" {
		logger.Error("DeleteMember", "error", "cannot delete owner")
		return nil, fmt.Errorf("cannot delete owner")
	}

	err = uc.userRepo.DeleteUser(deleteUserID)
	if err != nil {
		logger.Error("DeleteMember", "func", "DeleteUser()", "error", err.Error())
		return nil, err
	}

	return response.NewDeleteMemberResponse()
}

func (uc *organizationUsecase) GetScreenDashboard(userID, orgID string) (*response.GetScreenDashboardResponse, error) {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("GetScreenDashboard", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	if user.Role != "member" {
		logger.Error("GetScreenDashboard", "error", "user is not member")
		return nil, fmt.Errorf("user is not member")
	}

	records, err := uc.imageDataRepo.FindByUserID(userID)
	if err != nil {
		logger.Error("GetScreenDashboard", "func", "FindByUserID()", "error", err.Error())
		return nil, err
	}

	user, err = uc.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("GetScreenDashboard", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	return response.NewGetScreenDashboardResponse(*user, records)
}

func (uc *organizationUsecase) GetScreenDashboardForAdmin(userID, orgID string) (*response.GetScreenDashboardForAdminResponse, error) {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("GetScreenDashboardForAdmin", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	if user.Role == "member" {
		logger.Error("GetScreenDashboardForAdmin", "error", "user is not admin")
		return nil, fmt.Errorf("user is not admin")
	}

	records, err := uc.imageDataRepo.FindByOrganizationID(orgID)
	if err != nil {
		logger.Error("GetScreenDashboardForAdmin", "func", "FindByOrganizationID()", "error", err.Error())
		return nil, err
	}

	users, err := uc.repo.GetMember(orgID)
	if err != nil {
		logger.Error("GetScreenDashboardForAdmin", "func", "GetMember()", "error", err.Error())
		return nil, err
	}

	return response.NewGetScreenDashboardForAdminResponse(users, records)
}
