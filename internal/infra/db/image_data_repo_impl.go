package db

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	jptime "inbody-ocr-backend/pkg/jp_time"

	"github.com/uptrace/bun"
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
	query := `INSERT INTO image_data (id, organization_id, user_id, measurement_date_id, weight, height, muscle_weight, fat_weight, fat_percent, body_water, protein, mineral, point, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := jptime.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	_, err := r.db.Exec(query, data.ID, data.OrganizationID, data.UserID, data.MeasurementDateID, data.Weight, data.Height, data.MuscleWeight, data.FatWeight, data.FatPercent, data.BodyWater, data.Protein, data.Mineral, data.Point, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *imageDataRepository) CreateDataWithTx(tx bun.Tx, data entity.ImageData) (*entity.ImageData, error) {
	query := `INSERT INTO image_data (id, organization_id, user_id, measurement_date_id, weight, height, muscle_weight, fat_weight, fat_percent, body_water, protein, mineral, point, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	now := jptime.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	_, err := tx.Exec(query, data.ID, data.OrganizationID, data.UserID, data.MeasurementDateID, data.Weight, data.Height, data.MuscleWeight, data.FatWeight, data.FatPercent, data.BodyWater, data.Protein, data.Mineral, data.Point, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *imageDataRepository) FindByUserID(userID string) ([]entity.ImageData, error) {
	var records []entity.ImageData

	query := `SELECT * FROM image_data WHERE user_id = ?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record entity.ImageData
		var createdAt, updatedAt string
		err := rows.Scan(&record.ID, &record.OrganizationID, &record.UserID, &record.MeasurementDateID, &record.Weight, &record.Height, &record.MuscleWeight, &record.FatWeight, &record.FatPercent, &record.BodyWater, &record.Protein, &record.Mineral, &record.Point, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		record.CreatedAt, err = jptime.ParseDateTime(createdAt)
		if err != nil {
			return nil, err
		}

		record.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func (r *imageDataRepository) FindByOrganizationID(orgID string) ([]entity.ImageData, error) {
	var records []entity.ImageData

	query := `SELECT * FROM image_data WHERE organization_id = ?`

	rows, err := r.db.Query(query, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record entity.ImageData
		var createdAt, updatedAt string
		err := rows.Scan(&record.ID, &record.OrganizationID, &record.UserID, &record.MeasurementDateID, &record.Weight, &record.Height, &record.MuscleWeight, &record.FatWeight, &record.FatPercent, &record.BodyWater, &record.Protein, &record.Mineral, &record.Point, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		record.CreatedAt, err = jptime.ParseDateTime(createdAt)
		if err != nil {
			return nil, err
		}

		record.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func (r *imageDataRepository) BeginTransaction() (bun.Tx, error) {
	return r.db.Begin()
}
