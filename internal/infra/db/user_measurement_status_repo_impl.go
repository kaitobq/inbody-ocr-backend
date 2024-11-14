package db

import (
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
)

type userMeasurementStatusRepository struct {
	db *database.DB
}

func NewUserMeasurementStatusRepository(db *database.DB) repository.UserMeasurementStatusRepository {
	return &userMeasurementStatusRepository{
		db: db,
	}
}
