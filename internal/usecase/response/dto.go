package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"time"
)

type UserImageData struct {
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Weight       float64   `json:"weight"`
	Height       float64   `json:"height"`
	MuscleWeight float64   `json:"muscle_weight"`
	FatWeight    float64   `json:"fat_weight"`
	FatPercent   float64   `json:"fat_percent"`
	BodyWater    float64   `json:"body_water"`
	Protein      float64   `json:"protein"`
	Mineral      float64   `json:"mineral"`
	Point        uint      `json:"point"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUserImageData(user entity.User, data entity.ImageData) UserImageData {
	if user.ID != data.UserID {
		return UserImageData{}
	}

	return UserImageData{
		User: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   user.ID,
			Name: user.Name,
		},
		Weight:       data.Weight,
		Height:       data.Height,
		MuscleWeight: data.MuscleWeight,
		FatWeight:    data.FatWeight,
		FatPercent:   data.FatPercent,
		BodyWater:    data.BodyWater,
		Protein:      data.Protein,
		Mineral:      data.Mineral,
		Point:        data.Point,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func NewUserImageDataList(users []entity.User, data []entity.ImageData) []UserImageData {
	userDataMap := make(map[string]entity.User, len(users))
	for _, user := range users {
		userDataMap[user.ID] = user
	}

	userDataList := make([]UserImageData, 0, len(data))
	for _, d := range data {
		if user, ok := userDataMap[d.UserID]; ok {
			userDataList = append(userDataList, NewUserImageData(user, d))
		}
	}

	return userDataList
}

type MeasurementDateResponse struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMeasurementDateResponse(data entity.MeasurementDate) MeasurementDateResponse {
	return MeasurementDateResponse{
		ID:        data.ID,
		Date:      data.Date,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func NewMeasurementDateResponseList(data []entity.MeasurementDate) []MeasurementDateResponse {
	res := make([]MeasurementDateResponse, 0, len(data))
	for _, d := range data {
		res = append(res, NewMeasurementDateResponse(d))
	}
	return res
}
