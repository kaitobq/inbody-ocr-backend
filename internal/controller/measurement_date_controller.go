package controller

import (
	"inbody-ocr-backend/internal/controller/render"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeasurementDateController struct {
	uc usecase.MeasurementDateUsecase
}

func NewMeasurementDateController(uc usecase.MeasurementDateUsecase) *MeasurementDateController {
	return &MeasurementDateController{
		uc: uc,
	}
}

func (ct *MeasurementDateController) GetMeasurementDate(c *gin.Context) {
	user := xcontext.User(c)
	orgID := user.OrganizationID

	res, err := ct.uc.GetMeasurementDate(orgID)
	if err != nil {
		logging.Errorf(c, "GetMeasurementDate GetMeasurementDate %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(c, res)
}

func (ct *MeasurementDateController) CreateMeasurementDate(c *gin.Context) {
	req, err := request.NewCreateMeasurementDateRequest(c)
	if err != nil {
		logging.Errorf(c, "CreateMeasurementDate NewCreateMeasurementDateRequest %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusBadRequest)
		return
	}

	user := xcontext.User(c)

	res, err := ct.uc.CreateMeasurementDate(user, req.Date)
	if err != nil {
		logging.Errorf(c, "CreateMeasurementDate CreateMeasurementDate %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(c, res)
}
