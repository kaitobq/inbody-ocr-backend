package usecase

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/usecase/response"
	"mime/multipart"
)

type UserUsecase interface {
	CreateUser(name, email, password, orgID string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
	GetOwnInfo(user entity.User) (*response.GetOwnInfoResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(userName, email, password, OrgName string) (*response.CreateOrganizationResponse, error)
	GetAllMembers(orgID string) (*response.GetAllMembersResponse, error)
	UpdateRole(updateUserID string, role entity.OrganizationRole, requestUser *entity.User) (*response.UpdateRoleResponse, error)
	DeleteMember(deleteUserID string, requestUser *entity.User) error
}

type ImageUsecase interface {
	AnalyzeImage(file multipart.File, fileHeader *multipart.FileHeader, user *entity.User) (*response.AnalyzeImageResponse, error)
}

type ImageDataUsecase interface {
	CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, user *entity.User, dateID string) error
	GetStatsForMember(user *entity.User) (*response.GetStatsForMemberResponse, error)
	GetStatsForAdmin(user *entity.User) (*response.GetStatsForAdminResponse, error)
	GetChartDataForMember(user *entity.User) (*response.GetChartDataForMemberResponse, error)
	GetChartDataForAdmin(user *entity.User) (*response.GetChartDataForAdminResponse, error)
	GetDataForMember(user *entity.User) (*response.GetImageDataForMemberResponse, error)
	GetDataForAdmin(user *entity.User) (*response.GetImageDataForAdminResponse, error)
	GetCurrentDataForAdmin(user *entity.User) (*response.GetCurrentImageDataForAdminResponse, error)
}

type MeasurementDateUsecase interface {
	GetMeasurementDate(orgID string) (*response.GetMeasurementDateResponse, error)
	CreateMeasurementDate(user *entity.User, dateStr string) (*response.CreateMeasurementDateResponse, error)
}
