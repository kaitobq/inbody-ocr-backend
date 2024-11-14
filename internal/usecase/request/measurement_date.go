package request

import "github.com/gin-gonic/gin"

type CreateMeasurementDateRequest struct {
	Date string `json:"date"`
}

func NewCreateMeasurementDateRequest(c *gin.Context) (*CreateMeasurementDateRequest, error) {
	var req CreateMeasurementDateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
