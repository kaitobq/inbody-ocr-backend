package repository

import "inbody-ocr-backend/internal/domain/entity"

type MeasurementDateRepository interface {
	FindByOrganizationID(orgID string) ([]entity.MeasurementDate, error)
	CreateMeasurementDate(date entity.MeasurementDate) error
}
