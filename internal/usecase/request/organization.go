package request

import "github.com/gin-gonic/gin"

type CreateOrganizationRequest struct {
	UserName     string `json:"user_name" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	OrgName string `json:"organization_name" binding:"required,min=3,max=32"`
}

func NewCreateOrganizationRequest(c *gin.Context) (*CreateOrganizationRequest, error) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
