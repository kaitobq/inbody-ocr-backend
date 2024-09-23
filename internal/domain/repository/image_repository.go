package repository

type ImageRepository interface {
	DetectTextFromImage(filePath, language string) ([]string, error)
}