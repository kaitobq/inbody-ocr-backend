package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logging"
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
		logging.Errorf(c, "CreateOrganization NewCreateOrganizationRequest %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.CreateOrganization(req.UserName, req.Email, req.Password, req.OrgName)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			logging.Errorf(c, "CreateOrganization CreateUser err={email already exists} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		default:
			logging.Errorf(c, "CreateOrganization CreateUser %v", err)
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	c.JSON(http.StatusCreated, res)
}

func (ct *OrganizationController) GetAllMembers(c *gin.Context) {
	_, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetAllMembers ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.GetAllMembers(orgID)
	if err != nil {
		logging.Errorf(c, "GetAllMembers GetAllMembers %v", err)
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *OrganizationController) UpdateRole(c *gin.Context) {
	updateUserID := c.Query("user_id")
	req, err := request.NewUpdateRoleRequest(c)
	if err != nil {
		logging.Errorf(c, "UpdateRole NewUpdateRoleRequest %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "UpdateRole ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.UpdateRole(updateUserID, req.Role, orgID, userID)
	if err != nil {
		switch err.Error() {
		case "user is not admin":
			logging.Errorf(c, "UpdateRole UpdateRole err={user is not admin} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		case "cannot update owner role":
			logging.Errorf(c, "UpdateRole UpdateRole err={cannot update owner role} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		default:
			logging.Errorf(c, "UpdateRole UpdateRole %v", err)
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, res)
}

func (ct *OrganizationController) DeleteMember(c *gin.Context) {
	deleteUserID := c.Query("user_id")

	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "DeleteMember ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.DeleteMember(deleteUserID, orgID, userID)
	if err != nil {
		switch err.Error() {
		case "user is not admin":
			logging.Errorf(c, "DeleteMember DeleteMember err={user is not admin} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		case "cannot delete owner":
			logging.Errorf(c, "DeleteMember DeleteMember err={cannot delete owner} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, res)
}
