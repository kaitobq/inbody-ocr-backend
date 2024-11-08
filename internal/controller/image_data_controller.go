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

type ImageDataController struct {
	uc           usecase.ImageDataUsecase
	tokenService service.TokenService
}

func NewImageDataController(uc usecase.ImageDataUsecase, tokenService service.TokenService) *ImageDataController {
	return &ImageDataController{
		uc:           uc,
		tokenService: tokenService,
	}
}

func (ct *ImageDataController) SaveImageData(c *gin.Context) {
	req, err := request.NewSaveImageDataRequest(c)
	if err != nil {
		logging.Errorf(c, "SaveImageData NewSaveImageDataRequest %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user := xcontext.MemberUser(c)

	res, err := ct.uc.CreateData(req.Weight, req.Height, req.MuscleWeight, req.FatWeight, req.FatPercent, req.BodyWater, req.Protein, req.Mineral, req.Point, user)
	if err != nil {
		logging.Errorf(c, "SaveImageData CreateData %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForMember(c *gin.Context) {
	user := xcontext.MemberUser(c)

	res, err := ct.uc.GetStatsForMember(user)
	if err != nil {
		logging.Errorf(c, "GetStatsForMember GetStatsForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForAdmin(c *gin.Context) {
	user := xcontext.AdminUser(c)

	res, err := ct.uc.GetStatsForAdmin(user)
	if err != nil {
		logging.Errorf(c, "GetStatsForAdmin GetStatsForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetChartDataForMember(c *gin.Context) {
	user := xcontext.MemberUser(c)

	res, err := ct.uc.GetChartDataForMember(user)
	if err != nil {
		logging.Errorf(c, "GetChartDataForMember GetChartDataForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetChartDataForAdmin(c *gin.Context) {
	user := xcontext.AdminUser(c)

	res, err := ct.uc.GetChartDataForAdmin(user)
	if err != nil {
		logging.Errorf(c, "GetChartDataForAdmin GetChartDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForMember(c *gin.Context) {
	user := xcontext.MemberUser(c)

	res, err := ct.uc.GetDataForMember(user)
	if err != nil {
		logging.Errorf(c, "GetImageDataForMember GetDataForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForAdmin(c *gin.Context) {
	user := xcontext.AdminUser(c)

	res, err := ct.uc.GetDataForAdmin(user)
	if err != nil {
		logging.Errorf(c, "GetImageDataForAdmin GetDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetCurrentImageDataForAdmin(c *gin.Context) {
	user := xcontext.AdminUser(c)

	res, err := ct.uc.GetCurrentDataForAdmin(user)
	if err != nil {
		logging.Errorf(c, "GetCurrentImageDataForAdmin GetCurrentDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
