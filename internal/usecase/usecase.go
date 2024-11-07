package usecase

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/usecase/response"
	"mime/multipart"
)

type UserUsecase interface {
	CreateUser(name, email, password, orgID string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
	GetOwnInfo(userID string) (*response.GetOwnInfoResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(userName, email, password, OrgName string) (*response.CreateOrganizationResponse, error)
	GetAllMembers(orgID string) (*response.GetAllMembersResponse, error)
	UpdateRole(updateUserID string, role entity.OrganizationRole, orgID, requestUserID string) (*response.UpdateRoleResponse, error)
	DeleteMember(deleteUserID, orgID, requestUserID string) (*response.DeleteMemberResponse, error)
	GetScreenDashboard(userID, orgID string) (*response.GetScreenDashboardResponse, error)
	GetScreenDashboardForAdmin(userID, orgID string) (*response.GetScreenDashboardForAdminResponse, error)
}

type ImageUsecase interface {
	AnalyzeImage(file multipart.File, fileHeader *multipart.FileHeader, userID, orgID string) (*response.AnalyzeImageResponse, error)
}

type ImageDataUsecase interface {
	CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, userID, orgID string) (*response.SaveImageDataResponse, error)
	GetStatsForMember(userID, orgID string) (*response.GetStatsForMemberResponse, error)
	GetDataForMember(userID string) (*response.GetImageDataForMemberResponse, error)
	GetDataForAdmin(userID, orgID string) (*response.GetImageDataForAdminResponse, error)
}
