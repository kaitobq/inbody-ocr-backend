package controller

import "inbody-ocr-backend/internal/usecase"

type ImageDataController struct {
	uc usecase.ImageDataUsecase
}

func NewImageDataController(uc usecase.ImageDataUsecase) *ImageDataController {
	return &ImageDataController{
		uc: uc,
	}
}

