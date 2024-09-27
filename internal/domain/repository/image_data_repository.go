package repository

import "inbody-ocr-backend/internal/domain/entity"

type ImageDataRepository interface {
	CreateData(data entity.ImageData) (*entity.ImageData, error)
	FindByUserID(userID string) ([]entity.ImageData, error)
	FindByOrganizationID(orgID string) ([]entity.ImageData, error)
}
