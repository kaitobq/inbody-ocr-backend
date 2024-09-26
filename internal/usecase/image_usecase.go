package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
	"inbody-ocr-backend/internal/usecase/response"
	"io"
	"mime/multipart"
	"os"
)

type imageUsecase struct {
	repo          repository.ImageRepository
	imageDataRepo repository.ImageDataRepository
	ulidService   service.ULIDService
}

func NewImageUsecase(repo repository.ImageRepository, ulidService service.ULIDService, imageDataRepo repository.ImageDataRepository) ImageUsecase {
	return &imageUsecase{
		repo:          repo,
		imageDataRepo: imageDataRepo,
		ulidService:   ulidService,
	}
}

func (uc *imageUsecase) AnalyzeImage(file multipart.File, userID, orgID string) (*response.AnalyzeImageResponse, error) {
	// 一時ファイルを作成して、画像データを保存
	tempFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		logger.Error("AnalyzeImage", "func", "CreateTemp()", "error", err.Error())
		return nil, err
	}
	defer os.Remove(tempFile.Name()) // 処理後、ファイルを削除
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		logger.Error("AnalyzeImage", "func", "io.Copy()", "error", err.Error())
		return nil, err
	}

	data, err := uc.repo.DetectTextFromImage(tempFile.Name(), "ja")
	if err != nil {
		logger.Error("AnalyzeImage", "func", "DetectTextFromImage()", "error", err.Error())
		return nil, err
	}

	// 実際はフロントから保存するリクエストを送らせるのでここで保存する必要はない
	id := uc.ulidService.GenerateULID()
	data.ID = id
	data.UserID = userID
	data.OrganizationID = orgID
	// data, err = uc.imageDataRepo.CreateData(*data)
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Printf("weight: %v\n", data.Weight)
	fmt.Printf("height: %v\n", data.Height)
	fmt.Printf("bodyWater: %v\n", data.BodyWater)
	fmt.Printf("fat: %v\n", data.FatWeight)
	fmt.Printf("fatp: %v\n", data.FatPercent)
	fmt.Printf("mineral: %v\n", data.Mineral)
	fmt.Printf("pr: %v\n", data.Protein)
	fmt.Printf("mw: %v\n", data.MuscleWeight)
	fmt.Printf("point: %v\n", data.Point)

	return response.NewAnalyzeImageResponse(*data), nil
}
