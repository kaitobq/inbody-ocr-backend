package repository

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/uptrace/bun"
)

type ImageDataRepository interface {
	CreateData(data entity.ImageData) (*entity.ImageData, error)
	CreateDataWithTx(tx bun.Tx, data entity.ImageData) (*entity.ImageData, error)
	FindByUserID(userID string) ([]entity.ImageData, error)
	FindByOrganizationID(orgID string) ([]entity.ImageData, error)
	BeginTransaction() (bun.Tx, error)
}
