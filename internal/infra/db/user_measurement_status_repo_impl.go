package db

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	jptime "inbody-ocr-backend/pkg/jp_time"

	"github.com/uptrace/bun"
)

type userMeasurementStatusRepository struct {
	db *database.DB
}

func NewUserMeasurementStatusRepository(db *database.DB) repository.UserMeasurementStatusRepository {
	return &userMeasurementStatusRepository{
		db: db,
	}
}

func (r *userMeasurementStatusRepository) CreateUserMeasurementStatus(ent entity.UserMeasurementStatus) error {
	query := `INSERT INTO user_measurement_status (id, user_id, measurement_date_id, image_data_id, has_registered, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	now := jptime.Now()
	ent.CreatedAt = now
	ent.UpdatedAt = now

	_, err := r.db.Exec(query, ent.ID, ent.UserID, ent.MeasurementDateID, ent.ImageDataID, ent.HasRegistered, ent.CreatedAt, ent.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) BeginTransaction() (bun.Tx, error) {
	return r.db.Begin()
}

func (r *userMeasurementStatusRepository) CreateUserMeasurementStatusWithTx(tx bun.Tx, ent entity.UserMeasurementStatus) error {
	query := `INSERT INTO user_measurement_status (id, user_id, measurement_date_id, image_data_id, has_registered, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	now := jptime.Now()
	ent.CreatedAt = now
	ent.UpdatedAt = now

	_, err := tx.Exec(query, ent.ID, ent.UserID, ent.MeasurementDateID, ent.ImageDataID, ent.HasRegistered, ent.CreatedAt, ent.UpdatedAt)
	return err
}
