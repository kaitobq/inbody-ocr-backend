package usecase

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/infra/logger"
	"inbody-ocr-backend/internal/usecase/response"
)

type imageDataUsecase struct {
	repo repository.ImageDataRepository
}

func NewImageDataUsecase(repo repository.ImageDataRepository) ImageDataUsecase {
	return &imageDataUsecase{
		repo: repo,
	}
}

func (uc *imageDataUsecase) CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, userID, orgID string) (*response.SaveImageDataResponse, error) {
	imageData := &entity.ImageData{
		UserID:       userID,
		OrganizationID:        orgID,
		Weight:       weight,
		Height:       height,
		MuscleWeight: muscleWeight,
		FatWeight:    fatWeight,
		FatPercent:   fatPercent,
		BodyWater:    bodyWater,
		Protein:      protein,
		Mineral:      mineral,
		Point:        point,
	}
	
	_, err := uc.repo.CreateData(*imageData)
	if err != nil {
		logger.Error("CreateData", "func", "CreateData()", "error", err.Error())
		return nil, err
	}

	return response.NewSaveImageDataResponse()
}
