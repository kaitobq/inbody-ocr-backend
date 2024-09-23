package controller

import (
	"fmt"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	uc usecase.ImageUsecase
}

func NewImageController(uc usecase.ImageUsecase) *ImageController {
	return &ImageController{
		uc: uc,
	}
}

// AnalyzeImage detects text from an uploaded image using Google Vision API
func (ct *ImageController) AnalyzeImage(c *gin.Context) {
	file, err := request.GetImgFileFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image from request"})
		return
	}

	texts, err := ct.uc.AnalyzeImage(file)
	if err != nil {
		fmt.Printf("Failed to detect text from image: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to detect text from image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"texts": texts})
}
