package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
	"inbody-ocr-backend/internal/usecase/response"
)

type userUsecase struct {
	repo         repository.UserRepository
	tokenService service.TokenService
	ulidService service.ULIDService
}

func NewUserUsecase(repo repository.UserRepository, tokenService service.TokenService, ulidService service.ULIDService) UserUsecase {
	return &userUsecase{
		repo: repo,
		tokenService: tokenService,
		ulidService: ulidService,
	}
}

func (uc *userUsecase) CreateUser(name, email, password string) (*response.SignUpResponse, error) {
	exists, err := uc.repo.CheckDuplicateEmail(email)
	if err != nil {
		logger.Error("CreateUser", "func", "CheckDuplicateEmail()", "error", err.Error())
		return nil, err
	}
	if exists {
		logger.Error("CreateUser", "func", "CheckDuplicateEmail()", "error", "email already exists")
		return nil, fmt.Errorf("email already exists")
	}

	hashedPassword, err := uc.repo.HashPassword(password)
	if err != nil {
		logger.Error("CreateUser", "func", "HashPassword()", "error", err.Error())
		return nil, err
	}

	id := uc.ulidService.GenerateULID()
	user := entity.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = uc.repo.CreateUser(user)
	if err != nil {
		logger.Error("CreateUser", "func", "CreateUser()", "error", err.Error())
		return nil, err
	}

	token, err := uc.tokenService.GenerateTokenFromID(id)
	if err != nil {
		logger.Error("CreateUser", "func", "GenerateTokenFromID", "error", err.Error())
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		logger.Error("CreateUser", "func", "ExtractExpFromToken", "error", err.Error())
		return nil, err
	}

	return response.NewSignUpResponse(token, exp)
}

func (uc *userUsecase) SignIn(email, password string) (*response.SignInResponse, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		logger.Error("SignIn", "func", "FindByEmail()", "error", err.Error())
		return nil, err
	}

	err = uc.repo.ComparePassword(user.Password, password)
	if err != nil {
		logger.Error("SignIn", "func", "ComparePassword()", "error", err.Error())
		return nil, err
	}
	
	token, err := uc.tokenService.GenerateTokenFromID(user.ID)
	if err != nil {
		logger.Error("SignIn", "func", "GenerateTokenFromID", "error", err.Error())
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		logger.Error("SignIn", "func", "ExtractExpFromToken", "error", err.Error())
		return nil, err
	}

	return response.NewSignInResponse(token, exp)
}
