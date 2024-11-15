package repository

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/uptrace/bun"
)

type MeasurementDateRepository interface {
	FindByID(id string) (*entity.MeasurementDate, error)
	FindByOrganizationID(orgID string) ([]entity.MeasurementDate, error)
	CreateMeasurementDate(date entity.MeasurementDate) error
	CountByOrganizationID(orgID string) (int, error)
	BeginTransaction() (bun.Tx, error)
	CreateMeasurementDateWithTx(tx bun.Tx, date entity.MeasurementDate) error
}
