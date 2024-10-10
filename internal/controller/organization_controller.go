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
