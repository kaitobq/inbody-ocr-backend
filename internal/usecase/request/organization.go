package request

import "github.com/gin-gonic/gin"

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

func NewCreateOrganizationRequest(c *gin.Context) (*CreateOrganizationRequest, error) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
