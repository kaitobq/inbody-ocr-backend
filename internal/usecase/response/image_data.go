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

type Kilo struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatWeight    float64 `json:"fat_weight"`
	FatPercent   float64 `json:"fat_percent"`
	BodyWater    float64 `json:"body_water"`
	Protein      float64 `json:"protein"`
	Mineral      float64 `json:"mineral"`
	CreatedAt    string  `json:"created_at"`
}

type Percent struct {
	FatPercent float64 `json:"fat_percent"`
	CreatedAt  string  `json:"created_at"`
}

type Score struct {
	Point     uint   `json:"point"`
	CreatedAt string `json:"created_at"`
}

type ChartDataForMember struct {
	Kilo    []Kilo    `json:"kilo"`
	Percent []Percent `json:"percent"`
	Score   []Score   `json:"score"`
}

type GetChartDataForMemberResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Chart   ChartDataForMember `json:"chart"`
}

func NewGetChartDataForMemberResponse(kilo []Kilo, percent []Percent, score []Score) (*GetChartDataForMemberResponse, error) {
	return &GetChartDataForMemberResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Chart: ChartDataForMember{
			Kilo:    kilo,
			Percent: percent,
			Score:   score,
		},
	}, nil
}

type ChartDataForAdminMap map[string]int

type ChartDataForAdmin struct {
	BMI          ChartDataForAdminMap `json:"bmi"`
	Weight       ChartDataForAdminMap `json:"weight"`
	MuscleWeight ChartDataForAdminMap `json:"muscle_weight"`
	FatWeight    ChartDataForAdminMap `json:"fat_weight"`
}

type GetChartDataForAdminResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Chart   ChartDataForAdmin `json:"chart"`
}

func NewGetChartDataForAdminResponse(chart ChartDataForAdmin) (*GetChartDataForAdminResponse, error) {
	return &GetChartDataForAdminResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Chart:   chart,
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
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []UserImageData `json:"data"`
}

func NewGetImageDataForAdminResponse(data []UserImageData) (*GetImageDataForAdminResponse, error) {
	return &GetImageDataForAdminResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Data:    data,
	}, nil
}

type GetCurrentImageDataForAdminResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []UserImageData `json:"data"`
}

func NewGetCurrentImageDataForAdminResponse(data []UserImageData) (*GetCurrentImageDataForAdminResponse, error) {
	return &GetCurrentImageDataForAdminResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Data:    data,
	}, nil
}
