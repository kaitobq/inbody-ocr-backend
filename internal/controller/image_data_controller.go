package controller

import (
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
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
		logger.Error("SaveImageData", "func", "NewSaveImageDataRequest()", "error", err.Error())
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logger.Error("SaveImageData", "func", "ExtractIDsFromContext()", "error", err.Error())
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.CreateData(req.Weight, req.Height, req.MuscleWeight, req.FatWeight, req.FatPercent, req.BodyWater, req.Protein, req.Mineral, req.Point, userID, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForMember(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logger.Error("GetStatsForMember", "func", "ExtractIDsFromContext()", "error", err.Error())
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetStatsForMember(userID, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetStatsForAdmin(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logger.Error("GetStatsForAdmin", "func", "ExtractIDsFromContext()", "error", err.Error())
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetStatsForAdmin(userID, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForMember(c *gin.Context) {
	userID, _, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logger.Error("GetImageDataForMember", "func", "ExtractIDsFromContext()", "error", err.Error())
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetDataForMember(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ct *ImageDataController) GetImageDataForAdmin(c *gin.Context) {
	userID, orgID, err := ct.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logger.Error("GetImageDataForAdmin", "func", "ExtractIDsFromContext()", "error", err.Error())
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	res, err := ct.uc.GetDataForAdmin(userID, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
