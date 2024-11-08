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

	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "SaveImageData ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.CreateData(req.Weight, req.Height, req.MuscleWeight, req.FatWeight, req.FatPercent, req.BodyWater, req.Protein, req.Mineral, req.Point, userID, orgID)
	if err != nil {
		logging.Errorf(c, "SaveImageData CreateData %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForMember(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetStatsForMember ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetStatsForMember(userID, orgID)
	if err != nil {
		logging.Errorf(c, "GetStatsForMember GetStatsForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForAdmin(c *gin.Context) {
	_, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetStatsForAdmin ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetStatsForAdmin(orgID)
	if err != nil {
		logging.Errorf(c, "GetStatsForAdmin GetStatsForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetChartDataForMember(c *gin.Context) {
	userID, _, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetChartDataForMember ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetChartDataForMember(userID)
	if err != nil {
		logging.Errorf(c, "GetChartDataForMember GetChartDataForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetChartDataForAdmin(c *gin.Context) {
	_, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetChartDataForAdmin ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetChartDataForAdmin(orgID)
	if err != nil {
		logging.Errorf(c, "GetChartDataForAdmin GetChartDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForMember(c *gin.Context) {
	userID, _, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetImageDataForMember ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetDataForMember(userID)
	if err != nil {
		logging.Errorf(c, "GetImageDataForMember GetDataForMember %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForAdmin(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetImageDataForAdmin ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetDataForAdmin(userID, orgID)
	if err != nil {
		logging.Errorf(c, "GetImageDataForAdmin GetDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetCurrentImageDataForAdmin(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "GetCurrentImageDataForAdmin ExtractIDsFromContext %v", err)
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetCurrentDataForAdmin(userID, orgID)
	if err != nil {
		logging.Errorf(c, "GetCurrentImageDataForAdmin GetCurrentDataForAdmin %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
