package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"
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
		logging.Errorf(c, "SignUp NewSignUpRequest %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.CreateUser(req.Name, req.Email, req.Password, req.OrgID)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			logging.Errorf(c, "SignUp CreateUser err={email already exists} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		default:
			logging.Errorf(c, "SignUp CreateUser %v", err)
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	c.JSON(http.StatusCreated, res)
}

func (ct *UserController) SignIn(c *gin.Context) {
	req, err := request.NewSignInRequest(c)
	if err != nil {
		logging.Errorf(c, "SignIn NewSignInRequest %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := ct.uc.SignIn(req.Email, req.Password)
	if err != nil {
		logging.Errorf(c, "SignIn SignIn %v", err)
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *UserController) GetOwnInfo(c *gin.Context) {
	user := xcontext.User(c)

	res, err := ct.uc.GetOwnInfo(*user)
	if err != nil {
		logging.Errorf(c, "GetOwnInfo GetOwnInfo %v", err)
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
