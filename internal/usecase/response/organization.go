package response

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"math"
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
		Message: "ok",
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

type GetAllMembersResponse struct {
	Status   int           `json:"status"`
	Message  string        `json:"message"`
	Members []UserResponse `json:"members"`
}

func NewGetAllMembersResponse(users []entity.User) (*GetAllMembersResponse, error) {
	var members []UserResponse
	for _, user := range users {
		members = append(members, *NewUserResponse(user))
	}

	return &GetAllMembersResponse{
		Status:   http.StatusOK,
		Message:  "ok",
		Members: members,
	}, nil
}

type UpdateRoleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Updated UserResponse `json:"updated"`
}

func NewUpdateRoleResponse(user entity.User) (*UpdateRoleResponse, error) {
	return &UpdateRoleResponse{
		Status:  http.StatusOK,
		Message: "ok",
		Updated: *NewUserResponse(user),
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

type GetScreenDashboardResponse struct {
	Status   int             `json:"status"`
	Message  string          `json:"message"`
	Current  Current         `json:"current"`
	Previous Previous        `json:"previous"`
	Graph    Graph           `json:"graph"`
	History  []UserImageData `json:"history"`
}

func NewGetScreenDashboardResponse(user entity.User, data []entity.ImageData) (*GetScreenDashboardResponse, error) {
	status := http.StatusOK
	message := "Dashboard data fetched successfully"

	if len(data) == 0 {
		return &GetScreenDashboardResponse{
			Status:  http.StatusOK,
			Message: "No data available",
		}, nil
	}

	// current
	sort.Slice(data, func(i, j int) bool {
		return data[i].CreatedAt.After(data[j].CreatedAt)
	})
	current := Current{
		Weight:       data[0].Weight,
		MuscleWeight: data[0].MuscleWeight,
		FatWeight:    data[0].FatWeight,
		CreatedAt:    data[0].CreatedAt,
	}

	// previous
	var previous Previous
	if len(data) > 1 {
		previous = Previous{
			Weight:       data[1].Weight,
			MuscleWeight: data[1].MuscleWeight,
			FatWeight:    data[1].FatWeight,
			CreatedAt:    data[1].CreatedAt,
		}
	}

	// graph
	var kilo []Kilo
	var percent []Percent
	var score []Score
	for _, record := range data {
		kilo = append(kilo, Kilo{
			Weight:       record.Weight,
			MuscleWeight: record.MuscleWeight,
			FatWeight:    record.FatWeight,
			BodyWater:    record.BodyWater,
			Protein:      record.Protein,
			Mineral:      record.Mineral,
			CreatedAt:    record.CreatedAt,
		})
		percent = append(percent, Percent{
			FatPercent: record.FatPercent,
			CreatedAt:  record.CreatedAt,
		})
		score = append(score, Score{
			Point:     record.Point,
			CreatedAt: record.CreatedAt,
		})
	}

	// history
	history := NewUserImageDataList([]entity.User{user}, data)

	return &GetScreenDashboardResponse{
		Status:   status,
		Message:  message,
		Current:  current,
		Previous: previous,
		Graph: Graph{
			Kilo:    kilo,
			Percent: percent,
			Score:   score,
		},
		History: history,
	}, nil
}

type Avg struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatPercent   float64 `json:"fat_percent"`
	Point        uint    `json:"point"`
}

type Transition struct {
	Weight       float64 `json:"weight"`
	MuscleWeight float64 `json:"muscle_weight"`
	FatPercent   float64 `json:"fat_percent"`
	Point        uint    `json:"point"`
}

type Chart struct {
	BMI          map[string]int `json:"bmi"`
	Weight       map[string]int `json:"weight"`
	MuscleWeight map[string]int `json:"muscle_weight"`
	FatPercent   map[string]int `json:"fat_percent"`
}

type UserPoint struct {
	Rank      uint      `json:"rank"`
	Name      string    `json:"name"`
	Point     uint      `json:"point"`
	CreatedAt time.Time `json:"created_at"`
}

type GetScreenDashboardForAdminResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Avg     Avg             `json:"avg"`
	Chart   Chart           `json:"chart"`
	Current []UserImageData `json:"current"`
	AllData []UserImageData `json:"all_data"`
}

func NewGetScreenDashboardForAdminResponse(users []entity.User, data []entity.ImageData) (*GetScreenDashboardForAdminResponse, error) {
	if len(data) == 0 {
		return &GetScreenDashboardForAdminResponse{
			Status:  http.StatusOK,
			Message: "No data available",
		}, nil
	}

	fmt.Println("NewGetScreenDashboardForAdminResponse", users, data)
	status := http.StatusOK
	message := "Dashboard data fetched successfully"

	latestRecords := getLatestRecords(data)
	weight, BMI, fatPercent, muscleWeight, points := collectData(latestRecords)

	// avg
	weightAvg := calcAvg(weight)
	muscleWeightAvg := calcAvg(muscleWeight)
	fatPercentAvg := calcAvg(fatPercent)
	pointAvg := calcAvg(points)
	avg := Avg{
		Weight:       weightAvg,
		MuscleWeight: muscleWeightAvg,
		FatPercent:   fatPercentAvg,
		Point:        uint(pointAvg),
	}

	// chart
	weightMin, weightMax := findMinMax(weight)
	muscleWeightMin, muscleWeightMax := findMinMax(muscleWeight)
	fatPercentMin, fatPercentMax := findMinMax(fatPercent)
	bmiMin, bmiMax := findMinMax(BMI)

	weightBins := generateBins(weightMin, weightMax, 5)
	muscleWeightBins := generateBins(muscleWeightMin, muscleWeightMax, 5)
	fatPercentBins := generateBins(fatPercentMin, fatPercentMax, 5)
	bmiBins := generateBins(bmiMin, bmiMax, 2)

	weightCounts := binData(weight, weightBins, 5)
	muscleWeightCounts := binData(muscleWeight, muscleWeightBins, 5)
	fatPercentCounts := binData(fatPercent, fatPercentBins, 5)
	bmiCounts := binData(BMI, bmiBins, 2)

	chart := Chart{
		BMI:          bmiCounts,
		Weight:       weightCounts,
		MuscleWeight: muscleWeightCounts,
		FatPercent:   fatPercentCounts,
	}

	// current
	var temp []entity.ImageData
	for _, record := range latestRecords {
		temp = append(temp, record)
	}
	current := NewUserImageDataList(users, temp)

	allData := NewUserImageDataList(users, data)

	return &GetScreenDashboardForAdminResponse{
		Status:  status,
		Message: message,
		Avg:     avg,
		Chart:   chart,
		Current: current,
		AllData: allData,
	}, nil
}

func getLatestRecords(records []entity.ImageData) map[string]entity.ImageData {
	latestRecords := make(map[string]entity.ImageData)
	for _, record := range records {
		if existing, ok := latestRecords[record.UserID]; !ok || record.CreatedAt.After(existing.CreatedAt) {
			latestRecords[record.UserID] = record
		}
	}

	return latestRecords
}

func collectData(records map[string]entity.ImageData) (weights, BMIs, fatPercents, muscleWeights, points []float64) {
	for _, record := range records {
		weights = append(weights, record.Weight)
		BMI := record.Weight / ((record.Height / 100) * (record.Height / 100))
		BMIs = append(BMIs, BMI)
		fatPercents = append(fatPercents, record.FatPercent)
		muscleWeights = append(muscleWeights, record.MuscleWeight)
		points = append(points, float64(record.Point))
	}
	return
}

func calcAvg(data []float64) float64 {
	sum := 0.0
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

func findMinMax(data []float64) (float64, float64) {
	if len(data) == 0 {
		return 0, 0
	}

	min, max := data[0], data[0]
	for _, d := range data {
		if d < min {
			min = d
		}
		if d > max {
			max = d
		}
	}

	return min, max
}

func generateBins(min, max, binWidth float64) []float64 {
	bins := []float64{}
	start := binWidth * math.Floor(min/binWidth)
	end := binWidth * math.Ceil(max/binWidth)
	for b := start; b <= end; b += binWidth {
		bins = append(bins, b)
	}
	return bins
}

func binData(data []float64, bins []float64, binWidth float64) map[string]int {
	counts := make(map[string]int)
	for _, value := range data {
		found := false
		for i := 0; i < len(bins)-1; i++ {
			lower, upper := bins[i], bins[i+1]
			if value >= lower && value < upper {
				key := fmt.Sprintf("%.1f-%.1f", lower, upper)
				counts[key]++
				found = true
				break
			}
		}

		if !found && value >= bins[len(bins)-1] {
			lower, upper := bins[len(bins)-1], bins[len(bins)-1]+binWidth
			key := fmt.Sprintf("%.1f-%.1f", lower, upper)
			counts[key]++
		}
	}

	return counts
}
