package usecase

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase/response"
	jptime "inbody-ocr-backend/pkg/jp_time"
)

type measurementDateUsecase struct {
	repo        repository.MeasurementDateRepository
	ulidService service.ULIDService
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

func (uc *measurementDateUsecase) CreateMeasurementDate(user *entity.User, dateStr string) (*response.CreateMeasurementDateResponse, error) {
	measurementDate, err := jptime.ParseDate(dateStr)
	if err != nil {
		return nil, err
	}

	date := entity.MeasurementDate{
		ID:             uc.ulidService.GenerateULID(),
		OrganizationID: user.OrganizationID,
		Date:           measurementDate,
	}

	err = uc.repo.CreateMeasurementDate(date)
	if err != nil {
		return nil, err
	}

	return response.NewCreateMeasurementDateResponse(date)
}
