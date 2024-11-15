package entity

import "time"

type UserMeasurementStatus struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	MeasurementDateID string    `json:"measurement_date_id"`
	ImageDataID       *string   `json:"image_data_id"` // nil許容
	HasRegistered     bool      `json:"has_registered"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
