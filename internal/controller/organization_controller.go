package controller

import (
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
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ct.uc.CreateOrganization(req.UserName, req.Email, req.Password, req.OrgName)
	if err != nil {
		switch err {
		case xerror.ErrEmailAlreadyExists:
			logging.Errorf(c, "CreateOrganization CreateUser err={%v}", err)
			render.ErrorCodeJSON(c, err.Error(), http.StatusBadRequest, xerror.CodeEmailAlreadyExists)
			return
		default:
			logging.Errorf(c, "CreateOrganization CreateUser %v", err)
			render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	render.JSON(c, res)
}

func (ct *OrganizationController) GetAllMembers(c *gin.Context) {
	user := xcontext.AdminUser(c)

	res, err := ct.uc.GetAllMembers(user.OrganizationID)
	if err != nil {
		logging.Errorf(c, "GetAllMembers GetAllMembers %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(c, res)
}

func (ct *OrganizationController) UpdateRole(c *gin.Context) {
	updateUserID := c.Query("user_id")
	req, err := request.NewUpdateRoleRequest(c)
	if err != nil {
		logging.Errorf(c, "UpdateRole NewUpdateRoleRequest %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	user := xcontext.AdminUser(c)

	res, err := ct.uc.UpdateRole(updateUserID, req.Role, user)
	if err != nil {
		switch err {
		case xerror.ErrCannotUpdateOwnerRole:
			logging.Errorf(c, "UpdateRole UpdateRole err={%v}", err)
			render.ErrorCodeJSON(c, err.Error(), http.StatusBadRequest, xerror.CodeCannotUpdateOwnerRole)
			return
		default:
			logging.Errorf(c, "UpdateRole UpdateRole %v", err)
			render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	render.JSON(c, res)
}

func (ct *OrganizationController) DeleteMember(c *gin.Context) {
	deleteUserID := c.Query("user_id")
	user := xcontext.AdminUser(c)

	res, err := ct.uc.DeleteMember(deleteUserID, user)
	if err != nil {
		switch err {
		case xerror.ErrCannotDeleteOwner:
			logging.Errorf(c, "DeleteMember DeleteMember err={%v}", err)
			render.ErrorCodeJSON(c, err.Error(), http.StatusBadRequest, xerror.CodeCannotDeleteOwner)
			return
		default:
			render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	render.JSON(c, res)
}
