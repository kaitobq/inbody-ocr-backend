package main

import (
	"inbody-ocr-backend/internal/app"
	"inbody-ocr-backend/internal/infra/logger"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("initialize app failed: %v\n", err)
	}
	defer func() {
		if err := app.Close(); err != nil {
			log.Printf("close app failed: %v\n", err)
		}
	}()
	log.Println("app initialized")

	logger.Init()

	if err := app.Migrate(); err != nil {
		log.Printf("failed to migrate: %v\n", err)
	}
	log.Println("migrated")

	if err := app.Run(); err != nil {
		log.Printf("app downed: %v\n", err)
	}
}
