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
	"inbody-ocr-backend/internal/middleware"
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
	measurementDateRepository := db.NewMeasurementDateRepository(databaseDB)
	userMeasurementStatusRepository := db.NewUserMeasurementStatusRepository(databaseDB)
	tokenService := service.NewTokenService()
	ulidService := service.NewULIDService()
	userUsecase := usecase.NewUserUsecase(userRepository, organizationRepository, measurementDateRepository, userMeasurementStatusRepository, tokenService, ulidService)
	userController := controller.NewUserController(userUsecase, tokenService)
	imageDataRepository := db.NewImageDataRepository(databaseDB)
	organizationUsecase := usecase.NewOrganizationUsecase(organizationRepository, userRepository, imageDataRepository, tokenService, ulidService)
	organizationController := controller.NewOrganizationController(organizationUsecase, tokenService)
	imageRepository := vision_api.NewImageRepository()
	imageUsecase := usecase.NewImageUsecase(imageRepository, ulidService, imageDataRepository)
	imageController := controller.NewImageController(imageUsecase, tokenService)
	imageDataUsecase := usecase.NewImageDataUsecase(imageDataRepository, organizationRepository, measurementDateRepository, userMeasurementStatusRepository, ulidService)
	imageDataController := controller.NewImageDataController(imageDataUsecase, tokenService)
	measurementDateUsecase := usecase.NewMeasurementDateUsecase(measurementDateRepository, organizationRepository, userMeasurementStatusRepository, ulidService)
	measurementDateController := controller.NewMeasurementDateController(measurementDateUsecase)
	containerContainer := container.NewCtrl(userController, organizationController, imageController, imageDataController, measurementDateController)
	configConfig := config.New()
	middlewareMiddleware := middleware.NewMiddleware(tokenService, userRepository)
	app := container.NewApp(engine, containerContainer, configConfig, databaseDB, middlewareMiddleware)
	return app, nil
}

// wire.go:

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
