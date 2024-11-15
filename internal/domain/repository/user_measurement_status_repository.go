package repository

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/uptrace/bun"
)

type UserMeasurementStatusRepository interface {
	CreateUserMeasurementStatus(ent entity.UserMeasurementStatus) error
	BeginTransaction() (bun.Tx, error)
	CreateUserMeasurementStatusWithTx(tx bun.Tx, ent entity.UserMeasurementStatus) error
	FindByUserID(userID string) (*entity.UserMeasurementStatus, error)
	UpdateHasRegisteredByUserID(userID string, registered bool) error
	UpdateHasRegisteredByUserIDWithTx(tx bun.Tx, userID string, registered bool) error
	UpdateImageDataIDByUserID(userID string, imageDataID *string) error
	UpdateImageDataIDByUserIDWithTx(tx bun.Tx, userID string, imageDataID *string) error
	UpdateMeasurementDateIDByUserID(userID, measurementDateID string) error
	UpdateMeasurementDateIDByUserIDWithTx(tx bun.Tx, userID, measurementDateID string) error
}
