package request

import "github.com/gin-gonic/gin"

type SaveImageDataRequest struct {
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

func NewSaveImageDataRequest(c *gin.Context) (*SaveImageDataRequest, error) {
	var req SaveImageDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
