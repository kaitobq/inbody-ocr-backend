package request

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type CreateOrganizationRequest struct {
	UserName string `json:"user_name" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	OrgName  string `json:"organization_name" binding:"required,min=3,max=32"`
}

func NewCreateOrganizationRequest(c *gin.Context) (*CreateOrganizationRequest, error) {
	var req CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

type UpdateRoleRequest struct {
	Role entity.OrganizationRole `json:"role" binding:"required,oneof=owner admin member"`
}

func NewUpdateRoleRequest(c *gin.Context) (*UpdateRoleRequest, error) {
	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
