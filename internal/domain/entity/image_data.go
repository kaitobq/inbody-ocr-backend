package entity

import "time"

type ImageData struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	UserID         string    `json:"user_id"`
	Weight         float64   `json:"weight"`
	Height         float64   `json:"height"`
	MuscleWeight   float64   `json:"muscle_weight"`
	FatWeight      float64   `json:"fat_weight"`
	FatPercent     float64   `json:"fat_percent"`
	BodyWater      float64   `json:"body_water"`
	Protein        float64   `json:"protein"`
	Mineral        float64   `json:"mineral"`
	Point          uint      `json:"point"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
