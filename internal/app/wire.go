//go:build wireinject

package app

import (
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/internal/app/container"
	"inbody-ocr-backend/internal/controller"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/db"
	"inbody-ocr-backend/internal/usecase"
	"inbody-ocr-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func New() (*container.App, error) {
	wire.Build(
		provideGinEngine,
		config.New,
		config.NewDBConfig,
		database.New,

		// service
		service.NewTokenService,
		service.NewULIDService,

		container.NewApp,
		container.NewCtrl,

		// user
		controller.NewUserController,
		usecase.NewUserUsecase,
		db.NewUserRepository,
	)

	return nil, nil
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
