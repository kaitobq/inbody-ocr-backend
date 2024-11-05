package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"inbody-ocr-backend/internal/usecase/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	uc           usecase.OrganizationUsecase
	tokenService service.TokenService
}

func NewOrganizationController(uc usecase.OrganizationUsecase, tokenService service.TokenService) *OrganizationController {
	return &OrganizationController{
		uc:           uc,
		tokenService: tokenService,
	}
}

func (ct *OrganizationController) CreateOrganization(c *gin.Context) {
	req, err := request.NewCreateOrganizationRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.CreateOrganization(req.UserName, req.Email, req.Password, req.OrgName)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ct *OrganizationController) GetAllMembers(c *gin.Context) {
	_, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}
	
	res, err := ct.uc.GetAllMembers(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *OrganizationController) UpdateRole(c *gin.Context) {
	updateUserID := c.Query("user_id")
	req, err := request.NewUpdateRoleRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.UpdateRole(updateUserID, req.Role, orgID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *OrganizationController) GetScreenDashboard(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.GetScreenDashboard(userID, orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *OrganizationController) GetScreenDashboardForAdmin(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.GetScreenDashboardForAdmin(userID, orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
