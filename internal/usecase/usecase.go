package usecase

import (
	"inbody-ocr-backend/internal/usecase/response"
	"mime/multipart"
)

type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(name string, founderID string) (*response.CreateOrganizationResponse, error)
}

type UserOrganizationMembershipUsecase interface {
	DeleteMembership(executorID, deleteUserID, organizationID string) (*response.DeleteMembershipResponse, error)
}

type ImageUsecase interface {
	AnalyzeImage(file multipart.File, userID string) (*response.AnalyzeImageResponse, error)
}
