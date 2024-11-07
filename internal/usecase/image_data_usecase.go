package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
	"inbody-ocr-backend/internal/usecase/response"
)

type imageDataUsecase struct {
	repo             repository.ImageDataRepository
	organizationRepo repository.OrganizationRepository
	userRepo         repository.UserRepository
	ulidService      service.ULIDService
}

func NewImageDataUsecase(repo repository.ImageDataRepository, organizationRepo repository.OrganizationRepository, userRepo repository.UserRepository, ulidService service.ULIDService) ImageDataUsecase {
	return &imageDataUsecase{
		repo:             repo,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
		ulidService:      ulidService,
	}
}

func (uc *imageDataUsecase) CreateData(weight, height, muscleWeight, fatWeight, fatPercent, bodyWater, protein, mineral float64, point uint, userID, orgID string) (*response.SaveImageDataResponse, error) {
	imageData := &entity.ImageData{
		ID:             uc.ulidService.GenerateULID(),
		UserID:         userID,
		OrganizationID: orgID,
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
		logger.Error("CreateData", "func", "CreateData()", "error", err.Error())
		return nil, err
	}

	return response.NewSaveImageDataResponse()
}

func (uc *imageDataUsecase) GetStatsForMember(userID, orgID string) (*response.GetStatsForMemberResponse, error) {
	records, err := uc.repo.FindByUserID(userID)
	if err != nil {
		logger.Error("GetStatsForMember", "func", "FindByUserID()", "error", err.Error())
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

func (uc *imageDataUsecase) GetStatsForAdmin(userID, orgID string) (*response.GetStatsForAdminResponse, error) {
	records, err := uc.repo.FindByOrganizationID(orgID)
	if err != nil {
		logger.Error("GetStatsForAdmin", "func", "FindByOrganizationID()", "error", err.Error())
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

func (uc *imageDataUsecase) GetDataForMember(userID string) (*response.GetImageDataForMemberResponse, error) {
	records, err := uc.repo.FindByUserID(userID)
	if err != nil {
		logger.Error("GetDataForMember", "func", "FindByUserID()", "error", err.Error())
		return nil, err
	}

	return response.NewGetImageDataForMemberResponse(records)
}

func (uc *imageDataUsecase) GetDataForAdmin(userID, orgID string) (*response.GetImageDataForAdminResponse, error) {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		logger.Error("GetDataForAdmin", "func", "FindByID()", "error", err.Error())
		return nil, err
	}

	if user.Role != "admin" && user.Role != "owner" {
		logger.Error("GetDataForAdmin", "error", "user is not admin")
		return nil, fmt.Errorf("user is not admin")
	}

	records, err := uc.repo.FindByOrganizationID(orgID)
	if err != nil {
		logger.Error("GetDataForAdmin", "func", "FindByOrganizationID()", "error", err.Error())
		return nil, err
	}

	// Convert records slice to map[string][]entity.ImageData
	recordsMap := make(map[string][]entity.ImageData)
	for _, record := range records {
		recordsMap[record.UserID] = append(recordsMap[record.UserID], record)
	}

	members, err := uc.organizationRepo.GetMember(orgID)
	if err != nil {
		logger.Error("GetDataForAdmin", "func", "GetMember()", "error", err.Error())
		return nil, err
	}

	return response.NewGetImageDataForAdminResponse(recordsMap, members)
}
