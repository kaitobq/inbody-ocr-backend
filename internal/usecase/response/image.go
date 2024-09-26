package response

import "inbody-ocr-backend/internal/domain/entity"

type Results struct {
	Weight       float64 `json:"weight"`
	Height       float64 `json:"height"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatWeight    float64 `json:"fat_weight"`
	FatPercent   float64 `json:"fat_percent"`
	BodyWater    float64 `json:"body_water"`
	Protein      float64 `json:"protein"`
	Mineral      float64 `json:"mineral"`
	Point        uint    `json:"point"`
}

type AnalyzeImageResponse struct {
	Message string  `json:"message"`
	Results Results `json:"results"`
}

func NewAnalyzeImageResponse(results entity.ImageData) *AnalyzeImageResponse {
	return &AnalyzeImageResponse{
		Message: "Image analyzed successfully",
		Results: Results{
			Weight:       results.Weight,
			Height:       results.Height,
			MuscleWeight: results.MuscleWeight,
			FatWeight:    results.FatWeight,
			FatPercent:   results.FatPercent,
			BodyWater:    results.BodyWater,
			Protein:      results.Protein,
			Mineral:      results.Mineral,
			Point:        results.Point,
		},
	}
}
