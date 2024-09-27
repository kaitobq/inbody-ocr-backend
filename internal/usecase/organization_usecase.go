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
	repo         repository.OrganizationRepository
	userRepo     repository.UserRepository
	tokenService service.TokenService
	ulidService  service.ULIDService
}

func NewOrganizationUsecase(repo repository.OrganizationRepository, userRepo repository.UserRepository, tokenService service.TokenService, ulidService service.ULIDService) OrganizationUsecase {
	return &organizationUsecase{
		repo:         repo,
		userRepo:     userRepo,
		tokenService: tokenService,
		ulidService:  ulidService,
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
