package repository

import "inbody-ocr-backend/internal/domain/entity"

type ImageRepository interface {
	DetectTextFromImage(filePath, language string) (*entity.ImageData, error)
}
