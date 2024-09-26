package db

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	"time"
)

type imageDataRepository struct {
	db database.DB
}

func NewImageDataRepository(db *database.DB) repository.ImageDataRepository {
	return &imageDataRepository{
		db: *db,
	}
}

func (r *imageDataRepository) CreateData(data entity.ImageData) (*entity.ImageData, error) {
	query := `INSERT INTO image_data (id, organization_id, user_id, weight, height, muscle_weight, fat_weight, fat_percent, body_water, protein, mineral, point, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	_, err := r.db.Exec(query, data.ID, data.OrganizationID, data.UserID, data.Weight, data.Height, data.MuscleWeight, data.FatWeight, data.FatPercent, data.BodyWater, data.Protein, data.Mineral, data.Point, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
