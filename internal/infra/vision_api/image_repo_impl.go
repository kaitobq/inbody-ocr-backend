package vision_api

import (
	"context"
	"encoding/json"
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type imageRepository struct {
}

func NewImageRepository() repository.ImageRepository {
	return &imageRepository{}
}

// detectTextFromImage detects text in an image using Google Vision API
func (r *imageRepository) DetectTextFromImage(filePath, language string) (*entity.ImageData, error) {
	// Google Vision APIクライアントを作成
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// 画像ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Vision APIに送信するための画像を準備
	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	// テキストを検出
	annotations, err := client.DetectTexts(ctx, image, &visionpb.ImageContext{LanguageHints: []string{language}}, 10)
	if err != nil {
		return nil, err
	}

	// "体重"に最も近い数値を検索
	weight, err := findClosestNumber("体重", annotations)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '体重':", weight)
	}
	weightFloat, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	height, err := findClosestNumber("身長", annotations)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '身長':", height)
	}
	height = strings.ReplaceAll(height, "cm", "")
	heightFloat, err := strconv.ParseFloat(height, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	muscle, err := findClosestNumber("筋肉", annotations)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '筋肉':", muscle)
	}
	muscleFloat, err := strconv.ParseFloat(muscle, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	fat, err := findClosestNumber("脂肪", annotations)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '脂肪':", fat) // 除脂肪量が検出されているが、体脂肪量で検出するように修正したい
	}
	fatFloat, err := strconv.ParseFloat(fat, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}
	fatFloat = weightFloat - fatFloat

	mineral, err := findClosestNumberInDirection("ミネラル", annotations, "Y")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to 'ミネラル':", mineral)
	}
	mineralFloat, err := strconv.ParseFloat(mineral, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	protein, err := findClosestNumberInDirection("タンパク", annotations, "Y")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to 'タンパク':", protein)
	}
	proteinFloat, err := strconv.ParseFloat(protein, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	water, err := findClosestNumberInDirection("水分", annotations, "Y")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '水分':", water)
	}
	waterFloat, err := strconv.ParseFloat(water, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert weight to float: %v", err)
	}

	score, err := findClosestNumber("点数", annotations)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Closest number to '点数':", score)
	}
	s := strings.Split(score, "/")
	scoreUInt, _ := strconv.ParseUint(s[0], 10, 32)

	var imgData entity.ImageData
	// imgData.OrganizationID
	// imgData.UserID
	imgData.Weight = weightFloat
	imgData.Height = heightFloat
	imgData.MuscleWeight = muscleFloat
	imgData.FatWeight = fatFloat
	imgData.FatPercent = fatFloat / weightFloat * 100
	imgData.Mineral = mineralFloat
	imgData.Protein = proteinFloat
	imgData.BodyWater = waterFloat
	imgData.Point = uint(scoreUInt)

	// annotationsをJSON形式に変換
	annotationsJSON, err := json.Marshal(annotations)
	if err != nil {
		return nil, err
	}

	// JSONデータをログに出力
	fmt.Printf("Annotations JSON: %s\n", string(annotationsJSON))

	// 検出されたテキストを格納
	// var detectedTexts []string
	// for _, annotation := range annotations {
	// 	detectedTexts = append(detectedTexts, annotation.Description)
	// }

	return &imgData, nil
}

// findClosestNumber finds the number closest to the given keyword in the list of annotations.
func findClosestNumber(keyword string, annotations []*visionpb.EntityAnnotation) (string, error) {
	var keywordAnnotation *visionpb.EntityAnnotation
	var closestNumber string
	minDistance := math.MaxFloat64

	// Regular expression to match numbers (both integers and floats)
	numberRegex := regexp.MustCompile(`\d+(\.\d+)?`)

	// Find the annotation that matches the keyword
	for _, annotation := range annotations {
		if annotation.Description == keyword {
			keywordAnnotation = annotation
			break
		}
	}

	// If the keyword is not found, return an error
	if keywordAnnotation == nil {
		return "", fmt.Errorf("keyword '%s' not found in annotations", keyword)
	}

	// Get the center of the keyword's bounding box
	keywordCenterX, keywordCenterY := getCenter(keywordAnnotation.BoundingPoly)

	// Iterate through all annotations to find the closest number
	for _, annotation := range annotations {
		// Skip the keyword itself
		if annotation.Description == keyword {
			continue
		}

		// Check if the annotation is a number
		if numberRegex.MatchString(annotation.Description) {
			// Get the center of the current annotation's bounding box
			currentCenterX, currentCenterY := getCenter(annotation.BoundingPoly)

			// Calculate the distance between the keyword and the current annotation
			distance := calculateDistance(keywordCenterX, keywordCenterY, currentCenterX, currentCenterY)

			// If this distance is smaller than the previous minimum, update the closest number
			if distance < minDistance {
				minDistance = distance
				closestNumber = annotation.Description
			}
		}
	}

	// If no number is found, return an error
	if closestNumber == "" {
		return "", fmt.Errorf("no number found near the keyword '%s'", keyword)
	}

	return closestNumber, nil
}

// getCenter calculates the center of a bounding box.
func getCenter(boundingPoly *visionpb.BoundingPoly) (float64, float64) {
	xSum, ySum := 0, 0
	for _, vertex := range boundingPoly.Vertices {
		xSum += int(vertex.X)
		ySum += int(vertex.Y)
	}
	return float64(xSum) / float64(len(boundingPoly.Vertices)), float64(ySum) / float64(len(boundingPoly.Vertices))
}

// calculateDistance calculates the Euclidean distance between two points.
func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

//

// findClosestNumberInDirection finds the number closest to the given keyword in the list of annotations based on the specified direction (X or Y).
func findClosestNumberInDirection(keyword string, annotations []*visionpb.EntityAnnotation, direction string) (string, error) {
	var keywordAnnotation *visionpb.EntityAnnotation
	var closestNumber string
	minDistance := math.MaxFloat64

	// Regular expression to match numbers (both integers and floats)
	numberRegex := regexp.MustCompile(`\d+(\.\d+)?`)

	// Find the annotation that matches the keyword
	for _, annotation := range annotations {
		if annotation.Description == keyword {
			keywordAnnotation = annotation
			break
		}
	}

	// If the keyword is not found, return an error
	if keywordAnnotation == nil {
		return "", fmt.Errorf("keyword '%s' not found in annotations", keyword)
	}

	// Get the center of the keyword's bounding box
	keywordCenterX, keywordCenterY := getCenter(keywordAnnotation.BoundingPoly)

	// Iterate through all annotations to find the closest number in the specified direction
	for _, annotation := range annotations {
		// Skip the keyword itself
		if annotation.Description == keyword {
			continue
		}

		// Check if the annotation is a number
		if numberRegex.MatchString(annotation.Description) {
			// Get the center of the current annotation's bounding box
			currentCenterX, currentCenterY := getCenter(annotation.BoundingPoly)

			var distance float64
			// Calculate the distance based on the direction (X or Y)
			if direction == "X" {
				distance = math.Abs(currentCenterX - keywordCenterX)
			} else if direction == "Y" {
				distance = math.Abs(currentCenterY - keywordCenterY)
			} else {
				return "", fmt.Errorf("invalid direction '%s', must be 'X' or 'Y'", direction)
			}

			// If this distance is smaller than the previous minimum, update the closest number
			if distance < minDistance {
				minDistance = distance
				closestNumber = annotation.Description
			}
		}
	}

	// If no number is found, return an error
	if closestNumber == "" {
		return "", fmt.Errorf("no number found near the keyword '%s' in direction '%s'", keyword, direction)
	}

	return closestNumber, nil
}

// getClosestNumbersByDirection gets the closest numbers in both X and Y directions for the specified keyword.
// func getClosestNumbersByDirection(keyword string, annotations []*visionpb.EntityAnnotation) (string, string, error) {
// 	closestNumberX, errX := findClosestNumberInDirection(keyword, annotations, "X")
// 	if errX != nil {
// 		return "", "", errX
// 	}

// 	closestNumberY, errY := findClosestNumberInDirection(keyword, annotations, "Y")
// 	if errY != nil {
// 		return "", "", errY
// 	}

// 	return closestNumberX, closestNumberY, nil
// }
