package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"inbody-ocr-backend/internal/usecase/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc           usecase.UserUsecase
	tokenService service.TokenService
}

func NewUserController(uc usecase.UserUsecase, tokenService service.TokenService) *UserController {
	return &UserController{
		uc:           uc,
		tokenService: tokenService,
	}
}

func (ct *UserController) SignUp(c *gin.Context) {
	req, err := request.NewSignUpRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.CreateUser(req.Name, req.Email, req.Password, req.OrgID)
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

func (ct *UserController) SignIn(c *gin.Context) {
	req, err := request.NewSignInRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *UserController) Authenticate(c *gin.Context) {
	ok, err := ct.tokenService.TokenValid(c)
	if err != nil || !ok {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	userID, _, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	res, err := ct.uc.Authenticate(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
