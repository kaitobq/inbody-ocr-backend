package usecase

import (
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/usecase/response"
)

type measurementDateUsecase struct {
	repo repository.MeasurementDateRepository
}

func NewMeasurementDateUsecase(repo repository.MeasurementDateRepository) MeasurementDateUsecase {
	return &measurementDateUsecase{
		repo: repo,
	}
}

func (uc *measurementDateUsecase) GetMeasurementDate(orgID string) (*response.GetMeasurementDateResponse, error) {
	dates, err := uc.repo.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	return response.NewGetMeasurementDateResponse(dates)
}
