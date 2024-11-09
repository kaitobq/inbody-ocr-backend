package controller

import (
	"errors"
	"inbody-ocr-backend/internal/controller/render"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/domain/xerror"
	"inbody-ocr-backend/internal/infra/logging"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
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
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ct.uc.CreateUser(req.Name, req.Email, req.Password, req.OrgID)
	if err != nil {
		switch {
		case errors.Is(err, xerror.ErrEmailAlreadyExists):
			logging.Errorf(c, "SignUp CreateUser err={%v}", err)
			render.ErrorCodeJSON(c, err.Error(), http.StatusBadRequest, xerror.CodeEmailAlreadyExists)
			return
		default:
			logging.Errorf(c, "SignUp CreateUser %v", err)
			render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	render.JSON(c, res)
}

func (ct *UserController) SignIn(c *gin.Context) {
	req, err := request.NewSignInRequest(c)
	if err != nil {
		logging.Errorf(c, "SignIn NewSignInRequest %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ct.uc.SignIn(req.Email, req.Password)
	if err != nil {
		logging.Errorf(c, "SignIn SignIn %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(c, res)
}

func (ct *UserController) GetOwnInfo(c *gin.Context) {
	user := xcontext.User(c)

	res, err := ct.uc.GetOwnInfo(*user)
	if err != nil {
		logging.Errorf(c, "GetOwnInfo GetOwnInfo %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(c, res)
}
