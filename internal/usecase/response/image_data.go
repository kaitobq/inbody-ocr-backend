package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	jptime "inbody-ocr-backend/pkg/jp_time"
)

type StatsForMember struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatWeight    float64 `json:"fat_weight"`
	CreatedAt    string  `json:"created_at"`
}

type GetStatsForMemberResponse struct {
	Current  StatsForMember `json:"current"`
	Previous StatsForMember `json:"previous"`
}

func NewGetStatsForMemberResponse(current entity.ImageData, previous entity.ImageData) (*GetStatsForMemberResponse, error) {
	currentCreatedAt := jptime.FormatDateTime(current.CreatedAt)
	previousCreatedAt := jptime.FormatDateTime(previous.CreatedAt)

	return &GetStatsForMemberResponse{
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
	Stats StatsForAdmin `json:"stats"`
}

func NewGetStatsForAdminResponse(stats StatsForAdmin) (*GetStatsForAdminResponse, error) {
	return &GetStatsForAdminResponse{
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
	Chart ChartDataForMember `json:"chart"`
}

func NewGetChartDataForMemberResponse(kilo []Kilo, percent []Percent, score []Score) (*GetChartDataForMemberResponse, error) {
	return &GetChartDataForMemberResponse{
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
	Chart ChartDataForAdmin `json:"chart"`
}

func NewGetChartDataForAdminResponse(chart ChartDataForAdmin) (*GetChartDataForAdminResponse, error) {
	return &GetChartDataForAdminResponse{
		Chart: chart,
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
	Data []UserImageData `json:"data"`
}

func NewGetImageDataForAdminResponse(data []UserImageData) (*GetImageDataForAdminResponse, error) {
	return &GetImageDataForAdminResponse{
		Data: data,
	}, nil
}

type GetCurrentImageDataForAdminResponse struct {
	Data []UserImageData `json:"data"`
}

func NewGetCurrentImageDataForAdminResponse(data []UserImageData) (*GetCurrentImageDataForAdminResponse, error) {
	return &GetCurrentImageDataForAdminResponse{
		Data: data,
	}, nil
}
