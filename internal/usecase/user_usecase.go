package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase/response"
)

type userUsecase struct {
	repo         repository.UserRepository
	orgRepo      repository.OrganizationRepository
	tokenService service.TokenService
	ulidService  service.ULIDService
}

func NewUserUsecase(repo repository.UserRepository, orgRepo repository.OrganizationRepository, tokenService service.TokenService, ulidService service.ULIDService) UserUsecase {
	return &userUsecase{
		repo:         repo,
		orgRepo:      orgRepo,
		tokenService: tokenService,
		ulidService:  ulidService,
	}
}

// この関数はユーザが組織に参加するときに使用する
func (uc *userUsecase) CreateUser(name, email, password, orgID string) (*response.SignUpResponse, error) {
	exists, err := uc.repo.UserExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	hashedPassword, err := uc.repo.HashPassword(password)
	if err != nil {
		return nil, err
	}

	id := uc.ulidService.GenerateULID()
	user := &entity.User{
		ID:             id,
		Name:           name,
		Email:          email,
		Password:       hashedPassword,
		OrganizationID: orgID,
		Role:           entity.OrganizationRoleMember,
	}

	user, err = uc.repo.CreateUser(*user)
	if err != nil {
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

	return response.NewSignUpResponse(token, exp, *user)
}

func (uc *userUsecase) SignIn(email, password string) (*response.SignInResponse, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = uc.repo.ComparePassword(user.Password, password)
	if err != nil {
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

	return response.NewSignInResponse(token, exp, *user)
}

func (uc *userUsecase) GetOwnInfo(userID string) (*response.GetOwnInfoResponse, error) {
	user, err := uc.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return response.NewGetOwnInfoResponse(*user)
}
