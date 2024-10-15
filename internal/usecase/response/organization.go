package response

import (
	"inbody-ocr-backend/internal/domain/entity"
	"net/http"
	"sort"
	"time"
)

type OrganizationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrganizationResponse struct {
	Status       int                  `json:"status"`
	Message      string               `json:"message"`
	Organization OrganizationResponse `json:"organization"`
	User         UserResponse         `json:"user"`
	Token        TokenResponse        `json:"token"`
}

func NewCreateOrganizationResponse(org entity.Organization, token, userID, userName string, exp time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
		Status:  http.StatusCreated,
		Message: "Organization created successfully",
		Organization: OrganizationResponse{
			ID:        org.ID,
			Name:      org.Name,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		},
		User: UserResponse{
			ID:   userID,
			Name: userName,
			Role: entity.OrganizationRoleOwner,
		},
		Token: TokenResponse{
			Value: token,
			Exp:   exp,
		},
	}, nil
}

type Current struct {
	Weight       float64   `json:"weight"`
	MuscleWeight float64   `json:"muscle_weight"`
	FatWeight    float64   `json:"fat_weight"`
	CreatedAt    time.Time `json:"created_at"`
}

type Previous struct {
	Weight       float64   `json:"weight"`
	MuscleWeight float64   `json:"muscle_weight"`
	FatWeight    float64   `json:"fat_weight"`
	CreatedAt    time.Time `json:"created_at"`
}

type Kilo struct {
	Weight       float64   `json:"weight"`
	MuscleWeight float64   `json:"muscle_weight"`
	FatWeight    float64   `json:"fat_weight"`
	BodyWater    float64   `json:"body_water"`
	Protein      float64   `json:"protein"`
	Mineral      float64   `json:"mineral"`
	CreatedAt    time.Time `json:"created_at"`
}

type Percent struct {
	FatPercent float64   `json:"fat_percent"`
	CreatedAt  time.Time `json:"created_at"`
}

type Score struct {
	Point     uint      `json:"point"`
	CreatedAt time.Time `json:"created_at"`
}

type Graph struct {
	Kilo    []Kilo    `json:"kilo"`
	Percent []Percent `json:"percent"`
	Score   []Score   `json:"score"`
}

type GetScreenDashboard struct {
	Status   int                `json:"status"`
	Message  string             `json:"message"`
	Current  Current            `json:"current"`
	Previous Previous           `json:"previous"`
	Graph    Graph              `json:"graph"`
	History  []entity.ImageData `json:"history"`
}

func NewGetScreenDashboardResponse(data []entity.ImageData) (*GetScreenDashboard, error) {
	status := http.StatusOK
	message := "Dashboard data fetched successfully"

	sort.Slice(data, func(i, j int) bool {
		return data[i].CreatedAt.After(data[j].CreatedAt)
	})
	var current Current
	current.Weight = data[0].Weight
	current.MuscleWeight = data[0].MuscleWeight
	current.FatWeight = data[0].FatWeight
	current.CreatedAt = data[0].CreatedAt

	var previous Previous
	if len(data) > 1 {
		previous.Weight = data[1].Weight
		previous.MuscleWeight = data[1].MuscleWeight
		previous.FatWeight = data[1].FatWeight
		previous.CreatedAt = data[1].CreatedAt
	}

	var kilo []Kilo
	var percent []Percent
	var score []Score
	for _, d := range data {
		kilo = append(kilo, Kilo{
			Weight:       d.Weight,
			MuscleWeight: d.MuscleWeight,
			FatWeight:    d.FatWeight,
			BodyWater:    d.BodyWater,
			Protein:      d.Protein,
			Mineral:      d.Mineral,
			CreatedAt:    d.CreatedAt,
		})
		percent = append(percent, Percent{
			FatPercent: d.FatPercent,
			CreatedAt:  d.CreatedAt,
		})
		score = append(score, Score{
			Point:     d.Point,
			CreatedAt: d.CreatedAt,
		})
	}

	return &GetScreenDashboard{
		Status:   status,
		Message:  message,
		Current:  current,
		Previous: previous,
		Graph: Graph{
			Kilo:    kilo,
			Percent: percent,
			Score:   score,
		},
		History: data,
	}, nil
}
