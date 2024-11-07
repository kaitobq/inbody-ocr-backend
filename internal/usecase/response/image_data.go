package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	jptime "inbody-ocr-backend/pkg/jp_time"
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

type StatsForMember struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatWeight    float64 `json:"fat_weight"`
	CreatedAt    string  `json:"created_at"`
}

type GetStatsForMemberResponse struct {
	Status   int            `json:"status"`
	Message  string         `json:"message"`
	Current  StatsForMember `json:"current"`
	Previous StatsForMember `json:"previous"`
}

func NewGetStatsForMemberResponse(current entity.ImageData, previous entity.ImageData) (*GetStatsForMemberResponse, error) {
	currentCreatedAt := jptime.FormatDateTime(current.CreatedAt)
	previousCreatedAt := jptime.FormatDateTime(previous.CreatedAt)

	return &GetStatsForMemberResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Current: StatsForMember{
			Weight:       current.Weight,
			MuscleWeight: current.MuscleWeight,
			FatWeight:    current.FatWeight,
			CreatedAt:    currentCreatedAt,
		},
		Previous: StatsForMember{
			Weight:       previous.Weight,
			MuscleWeight: previous.MuscleWeight,
			FatWeight:    previous.FatWeight,
			CreatedAt:    previousCreatedAt,
		},
	}, nil
}

type StatsForAdmin struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatPercent   float64 `json:"fat_percent"`
	Point        uint    `json:"point"`
}

type GetStatsForAdminResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Stats   StatsForAdmin `json:"stats"`
}

func NewGetStatsForAdminResponse(stats StatsForAdmin) (*GetStatsForAdminResponse, error) {
	return &GetStatsForAdminResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Stats: StatsForAdmin{
			Weight:       stats.Weight,
			MuscleWeight: stats.MuscleWeight,
			FatPercent:   stats.FatPercent,
			Point:        stats.Point,
		},
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
