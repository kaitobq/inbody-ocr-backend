package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserOrganizationMembershipController struct {
	uc usecase.UserOrganizationMembershipUsecase
	tokenService service.TokenService
}

func NewUserOrganizationMembershipController(uc usecase.UserOrganizationMembershipUsecase, tokenService service.TokenService) *UserOrganizationMembershipController {
	return &UserOrganizationMembershipController{
		uc: uc,
		tokenService: tokenService,
	}
}

func (ct *UserOrganizationMembershipController) DeleteMembership(c *gin.Context) {
	req, err := request.NewDeleteMembershipRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	executorID, err := ct.tokenService.ExtractIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res, err := ct.uc.DeleteMembership(executorID, req.UserID, req.OrganizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
