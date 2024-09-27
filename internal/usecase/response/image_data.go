package response

import "inbody-ocr-backend/internal/domain/entity"

type SaveImageDataResponse struct {
	Message string `json:"message"`
}

func NewSaveImageDataResponse() (*SaveImageDataResponse, error) {
	return &SaveImageDataResponse{
		Message: "Data saved successfully",
	}, nil
}

type GetImageDataForMemberResponse struct {
	Records []entity.ImageData `json:"records"`
}

func NewGetImageDataForMemberResponse(records []entity.ImageData) (*GetImageDataForMemberResponse, error) {
	return &GetImageDataForMemberResponse{
		Records: records,
	}, nil
}

type GetImageDataForAdminResponse struct {
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
		Records: records,
		Users:   userResponses,
	}, nil
}
