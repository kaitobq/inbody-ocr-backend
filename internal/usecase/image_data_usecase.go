package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase/response"
	jptime "inbody-ocr-backend/pkg/jp_time"
	"math"
)

type imageDataUsecase struct {
	repo             repository.ImageDataRepository
	organizationRepo repository.OrganizationRepository
	ulidService      service.ULIDService
}

func NewImageDataUsecase(repo repository.ImageDataRepository, organizationRepo repository.OrganizationRepository, ulidService service.ULIDService) ImageDataUsecase {
	return &imageDataUsecase{
		repo:             repo,
		organizationRepo: organizationRepo,
		ulidService:      ulidService,
	}
}

func (uc *imageDataUsecase) CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, user *entity.User) (*response.SaveImageDataResponse, error) {
	imageData := &entity.ImageData{
		ID:             uc.ulidService.GenerateULID(),
		UserID:         user.ID,
		OrganizationID: user.OrganizationID,
		Weight:         weight,
		Height:         height,
		MuscleWeight:   muscleWeight,
		FatWeight:      fatWeight,
		FatPercent:     fatPercent,
		BodyWater:      bodyWater,
		Protein:        protein,
		Mineral:        mineral,
		Point:          point,
	}

	_, err := uc.repo.CreateData(*imageData)
	if err != nil {
		return nil, err
	}

	return response.NewSaveImageDataResponse()
}

func (uc *imageDataUsecase) GetStatsForMember(user *entity.User) (*response.GetStatsForMemberResponse, error) {
	records, err := uc.repo.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	current := entity.ImageData{}
	previous := entity.ImageData{}
	// Get the latest record
	for _, record := range records {
		if record.CreatedAt.After(current.CreatedAt) {
			current = record
		}
	}

	// Get the second latest record
	for _, record := range records {
		if record.CreatedAt.After(previous.CreatedAt) && record.CreatedAt != current.CreatedAt {
			previous = record
		}
	}

	return response.NewGetStatsForMemberResponse(current, previous)
}

func (uc *imageDataUsecase) GetStatsForAdmin(user *entity.User) (*response.GetStatsForAdminResponse, error) {
	records, err := uc.repo.FindByOrganizationID(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	latestRecords := getLatestRecords(records)
	stats := calcAvg(latestRecords)

	return response.NewGetStatsForAdminResponse(stats)
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

func calcAvg(latestRecords map[string]entity.ImageData) response.StatsForAdmin {
	var weight, muscleWeight, fatPercent float64
	var point uint
	for _, record := range latestRecords {
		weight += record.Weight
		muscleWeight += record.MuscleWeight
		fatPercent += record.FatPercent
		point += record.Point
	}

	weight /= float64(len(latestRecords))
	muscleWeight /= float64(len(latestRecords))
	fatPercent /= float64(len(latestRecords))
	point /= uint(len(latestRecords))

	return response.StatsForAdmin{
		Weight:       weight,
		MuscleWeight: muscleWeight,
		FatPercent:   fatPercent,
		Point:        point,
	}
}

func (uc *imageDataUsecase) GetChartDataForMember(user *entity.User) (*response.GetChartDataForMemberResponse, error) {
	records, err := uc.repo.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	// sort records by created_at
	records = sortRecords(records)

	kilo := make([]response.Kilo, 0)
	percent := make([]response.Percent, 0)
	score := make([]response.Score, 0)

	for _, record := range records {
		createdAt := jptime.FormatDateTime(record.CreatedAt)
		kilo = append(kilo, response.Kilo{
			Weight:       record.Weight,
			MuscleWeight: record.MuscleWeight,
			FatWeight:    record.FatWeight,
			FatPercent:   record.FatPercent,
			BodyWater:    record.BodyWater,
			Protein:      record.Protein,
			Mineral:      record.Mineral,
			CreatedAt:    createdAt,
		})

		percent = append(percent, response.Percent{
			FatPercent: record.FatPercent,
			CreatedAt:  createdAt,
		})

		score = append(score, response.Score{
			Point:     record.Point,
			CreatedAt: createdAt,
		})
	}

	return response.NewGetChartDataForMemberResponse(kilo, percent, score)
}

func sortRecords(records []entity.ImageData) []entity.ImageData {
	for i := 0; i < len(records); i++ {
		for j := i + 1; j < len(records); j++ {
			if records[i].CreatedAt.After(records[j].CreatedAt) {
				records[i], records[j] = records[j], records[i]
			}
		}
	}

	return records
}

func (uc *imageDataUsecase) GetChartDataForAdmin(user *entity.User) (*response.GetChartDataForAdminResponse, error) {
	records, err := uc.repo.FindByOrganizationID(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	records = sortRecords(records)

	var bmi response.ChartDataForAdminMap
	var weight response.ChartDataForAdminMap
	var muscleWeight response.ChartDataForAdminMap
	var fatWeight response.ChartDataForAdminMap

	latestRecords := getLatestRecords(records)

	weights, BMIs, fatPercents, muscleWeights := collectData(latestRecords)

	minWeight, maxWeight := findMinMax(weights)
	minBMI, maxBMI := findMinMax(BMIs)
	minFatPercent, maxFatPercent := findMinMax(fatPercents)
	minMuscleWeight, maxMuscleWeight := findMinMax(muscleWeights)

	weightBins := generateBins(minWeight, maxWeight, 5)
	bmiBins := generateBins(minBMI, maxBMI, 2)
	fatPercentBins := generateBins(minFatPercent, maxFatPercent, 5)
	muscleWeightBins := generateBins(minMuscleWeight, maxMuscleWeight, 5)

	weight = generateBinData(weights, weightBins, 5)
	bmi = generateBinData(BMIs, bmiBins, 2)
	fatWeight = generateBinData(fatPercents, fatPercentBins, 5)
	muscleWeight = generateBinData(muscleWeights, muscleWeightBins, 5)

	chart := response.ChartDataForAdmin{
		BMI:          bmi,
		Weight:       weight,
		MuscleWeight: muscleWeight,
		FatWeight:    fatWeight,
	}

	return response.NewGetChartDataForAdminResponse(chart)
}

func collectData(records map[string]entity.ImageData) (weights, BMIs, fatPercents, muscleWeights []float64) {
	for _, record := range records {
		weights = append(weights, record.Weight)
		BMI := record.Weight / ((record.Height / 100) * (record.Height / 100))
		BMIs = append(BMIs, BMI)
		fatPercents = append(fatPercents, record.FatPercent)
		muscleWeights = append(muscleWeights, record.MuscleWeight)
	}
	return
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

func generateBinData(data []float64, bins []float64, binWidth float64) map[string]int {
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

func (uc *imageDataUsecase) GetDataForMember(user *entity.User) (*response.GetImageDataForMemberResponse, error) {
	records, err := uc.repo.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return response.NewGetImageDataForMemberResponse(records)
}

func (uc *imageDataUsecase) GetDataForAdmin(user *entity.User) (*response.GetImageDataForAdminResponse, error) {
	records, err := uc.repo.FindByOrganizationID(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	users, err := uc.organizationRepo.GetMember(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	userImgData := response.NewUserImageDataList(users, records)

	return response.NewGetImageDataForAdminResponse(userImgData)
}

func (uc *imageDataUsecase) GetCurrentDataForAdmin(user *entity.User) (*response.GetCurrentImageDataForAdminResponse, error) {
	records, err := uc.repo.FindByOrganizationID(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	users, err := uc.organizationRepo.GetMember(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	latestRecords := getLatestRecords(records)

	var data []entity.ImageData
	for _, record := range latestRecords {
		data = append(data, record)
	}

	userImgData := response.NewUserImageDataList(users, data)

	return response.NewGetCurrentImageDataForAdminResponse(userImgData)
}
