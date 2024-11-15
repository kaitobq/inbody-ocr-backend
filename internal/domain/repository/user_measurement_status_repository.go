package repository

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/uptrace/bun"
)

type UserMeasurementStatusRepository interface {
	CreateUserMeasurementStatus(ent entity.UserMeasurementStatus) error
	BeginTransaction() (bun.Tx, error)
	CreateUserMeasurementStatusWithTx(tx bun.Tx, ent entity.UserMeasurementStatus) error
}
