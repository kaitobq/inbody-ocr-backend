package repository

import "inbody-ocr-backend/internal/domain/entity"

type ImageDataRepository interface {
	CreateData(data entity.ImageData) (*entity.ImageData, error)
}
