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
	user := entity.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := uc.tokenService.GenerateTokenFromID(id)
	if err != nil {
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewSignUpResponse(token, exp)
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
	
	token, err := uc.tokenService.GenerateTokenFromID(user.ID)
	if err != nil {
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewSignInResponse(token, exp)
}
