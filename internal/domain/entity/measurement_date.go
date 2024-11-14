package entity

import "time"

type MeasurementDate struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Date           string    `json:"date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
