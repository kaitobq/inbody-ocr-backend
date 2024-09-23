package usecase

import (
	"inbody-ocr-backend/internal/domain/repository"
	"io"
	"mime/multipart"
	"os"
)

type imageUsecase struct {
	repo repository.ImageRepository
}

func NewImageUsecase(repo repository.ImageRepository) ImageUsecase {
	return &imageUsecase{
		repo: repo,
	}
}

func (uc *imageUsecase) AnalyzeImage(file multipart.File) ([]string, error) {
	// 一時ファイルを作成して、画像データを保存
	tempFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name()) // 処理後、ファイルを削除
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	texts, err := uc.repo.DetectTextFromImage(tempFile.Name(), "ja")
	if err != nil {
		return nil, err
	}

	return texts, nil
}
