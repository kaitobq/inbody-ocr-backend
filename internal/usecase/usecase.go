package usecase

import (
	"inbody-ocr-backend/internal/usecase/response"
	"mime/multipart"
)

type UserUsecase interface {
	CreateUser(name, email, password, orgID string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(userName, email, password, OrgName string) (*response.CreateOrganizationResponse, error)
}

type ImageUsecase interface {
	AnalyzeImage(file multipart.File, userID, orgID string) (*response.AnalyzeImageResponse, error)
}
