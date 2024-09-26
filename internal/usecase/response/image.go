package response

import "inbody-ocr-backend/internal/domain/entity"

type AnalyzeImageResponse struct {
	Results entity.ImageData `json:"results"`
}

func NewAnalyzeImageResponse(results entity.ImageData) *AnalyzeImageResponse {
	return &AnalyzeImageResponse{
		Results: results,
	}
}
