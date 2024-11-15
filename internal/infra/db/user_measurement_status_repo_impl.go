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

func (r *userMeasurementStatusRepository) FindByUserID(userID string) (*entity.UserMeasurementStatus, error) {
	query := `SELECT * FROM user_measurement_status WHERE user_id = ?`

	row := r.db.QueryRow(query, userID)

	var createdAt, updatedAt string
	var ent entity.UserMeasurementStatus
	if err := row.Scan(&ent.ID, &ent.UserID, &ent.MeasurementDateID, &ent.ImageDataID, &ent.HasRegistered, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	var err error
	ent.CreatedAt, err = jptime.ParseDateTime(createdAt)
	if err != nil {
		return nil, err
	}
	ent.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (r *userMeasurementStatusRepository) UpdateHasRegisteredByUserID(userID string, registered bool) error {
	query := `UPDATE user_measurement_status SET has_registered = ? WHERE user_id = ?`

	_, err := r.db.Exec(query, true, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) UpdateHasRegisteredByUserIDWithTx(tx bun.Tx, userID string, registered bool) error {
	query := `UPDATE user_measurement_status SET has_registered = ? WHERE user_id = ?`

	_, err := tx.Exec(query, true, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) UpdateImageDataIDByUserID(userID string, imageDataID *string) error {
	query := `UPDATE user_measurement_status SET image_data_id = ? WHERE user_id = ?`

	_, err := r.db.Exec(query, imageDataID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) UpdateImageDataIDByUserIDWithTx(tx bun.Tx, userID string, imageDataID *string) error {
	query := `UPDATE user_measurement_status SET image_data_id = ? WHERE user_id = ?`

	_, err := tx.Exec(query, imageDataID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) UpdateMeasurementDateIDByUserID(userID string, measurementDateID string) error {
	query := `UPDATE user_measurement_status SET measurement_date_id = ? WHERE user_id = ?`

	_, err := r.db.Exec(query, measurementDateID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userMeasurementStatusRepository) UpdateMeasurementDateIDByUserIDWithTx(tx bun.Tx, userID string, measurementDateID string) error {
	query := `UPDATE user_measurement_status SET measurement_date_id = ? WHERE user_id = ?`

	_, err := tx.Exec(query, measurementDateID, userID)
	if err != nil {
		return err
	}

	return nil
}
