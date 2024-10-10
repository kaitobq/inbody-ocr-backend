package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"net/http"
)

type SaveImageDataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewSaveImageDataResponse() (*SaveImageDataResponse, error) {
	return &SaveImageDataResponse{
		Status:  http.StatusOK,
		Message: "Data saved successfully",
	}, nil
}

type GetImageDataForMemberResponse struct {
	Status  int                `json:"status"`
	Records []entity.ImageData `json:"records"`
}

func NewGetImageDataForMemberResponse(records []entity.ImageData) (*GetImageDataForMemberResponse, error) {
	return &GetImageDataForMemberResponse{
		Status:  http.StatusOK,
		Records: records,
	}, nil
}

type GetImageDataForAdminResponse struct {
	Status  int                           `json:"status"`
	Records map[string][]entity.ImageData `json:"records"`
	Users   []UserResponse                `json:"users"`
}

func NewGetImageDataForAdminResponse(records map[string][]entity.ImageData, users []entity.User) (*GetImageDataForAdminResponse, error) {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Role: user.Role,
		})
	}

	return &GetImageDataForAdminResponse{
		Status:  http.StatusOK,
		Records: records,
		Users:   userResponses,
	}, nil
}
