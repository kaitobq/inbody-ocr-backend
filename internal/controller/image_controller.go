package controller

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"inbody-ocr-backend/internal/usecase/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	uc           usecase.ImageUsecase
	tokenService service.TokenService
}

func NewImageController(uc usecase.ImageUsecase, tokenService service.TokenService) *ImageController {
	return &ImageController{
		uc:           uc,
		tokenService: tokenService,
	}
}

// AnalyzeImage detects text from an uploaded image using Google Vision API
func (ct *ImageController) AnalyzeImage(c *gin.Context) {
	file, fileHeader, err := request.GetImgFileFromContext(c)
	if err != nil {
		logging.Errorf(c, "AnalyzeImage GetImgFileFromContext %v", err)
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Failed to get image from request"))
		return
	}

	user := xcontext.MemberUser(c)

	res, err := ct.uc.AnalyzeImage(file, fileHeader, user)
	if err != nil {
		switch err.Error() {
		case "HEIC file is not supported":
			logging.Errorf(c, "AnalyzeImage AnalyzeImage err={HEIC file is not supported} %v", err)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		default:
			logging.Errorf(c, "AnalyzeImage AnalyzeImage %v", err)
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to detect text from image: %v", err)))
			return
		}
	}

	c.JSON(http.StatusOK, res)
}
