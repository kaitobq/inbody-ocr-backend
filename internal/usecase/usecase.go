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

type ImageDataUsecase interface {
	CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, userID, orgID string) (*response.SaveImageDataResponse, error)
	GetDataForMember(userID string) (*response.GetImageDataForMemberResponse, error)
	GetDataForAdmin(userID, orgID string) (*response.GetImageDataForAdminResponse, error)
}
