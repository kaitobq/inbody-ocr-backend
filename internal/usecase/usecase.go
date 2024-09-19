package usecase

import "inbody-ocr-backend/internal/usecase/response"

type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}
