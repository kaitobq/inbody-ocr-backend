// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/internal/app/container"
	"inbody-ocr-backend/internal/controller"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/db"
	"inbody-ocr-backend/internal/infra/vision_api"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/pkg/database"
)

// Injectors from wire.go:

func New() (*container.App, error) {
	engine := provideGinEngine()
	dbConfig := config.NewDBConfig()
	databaseDB, err := database.New(dbConfig)
	if err != nil {
		return nil, err
	}
	userRepository := db.NewUserRepository(databaseDB)
	organizationRepository := db.NewOrganizationRepository(databaseDB)
	tokenService := service.NewTokenService()
	ulidService := service.NewULIDService()
	userUsecase := usecase.NewUserUsecase(userRepository, organizationRepository, tokenService, ulidService)
	userController := controller.NewUserController(userUsecase)
	organizationUsecase := usecase.NewOrganizationUsecase(organizationRepository, userRepository, tokenService, ulidService)
	organizationController := controller.NewOrganizationController(organizationUsecase, tokenService)
	imageRepository := vision_api.NewImageRepository()
	imageDataRepository := db.NewImageDataRepository(databaseDB)
	imageUsecase := usecase.NewImageUsecase(imageRepository, ulidService, imageDataRepository)
	imageController := controller.NewImageController(imageUsecase, tokenService)
	containerContainer := container.NewCtrl(userController, organizationController, imageController, tokenService)
	configConfig := config.New()
	app := container.NewApp(engine, containerContainer, configConfig, databaseDB)
	return app, nil
}

// wire.go:

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
