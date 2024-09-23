package entity

type Image struct {
	ID        string `json:"id"`
	OrganizationID string `json:"organization_id"`
	UserID   string `json:"user_id"`
	Weight  float64 `json:"weight"`
	Height  float64 `json:"height"`
	MuscleWeight  float64 `json:"muscle_weight"`
	FatWeight  float64 `json:"fat_weight"`
	FatPercent float64 `json:"fat_percent"`
	Protein float64 `json:"protein"`
	Mineral float64 `json:"mineral"`
	Point   uint `json:"point"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
