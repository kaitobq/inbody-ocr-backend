package usecase

import "inbody-ocr-backend/internal/domain/repository"

type imageDataUsecase struct {
	repo repository.ImageDataRepository
}

func NewImageDataUsecase(repo repository.ImageDataRepository) ImageDataUsecase {
	return &imageDataUsecase{
		repo: repo,
	}
}
