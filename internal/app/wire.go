//go:build wireinject

package app

import (
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/internal/app/container"
	"inbody-ocr-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func New() (*container.App, error) {
	wire.Build(
		provideGinEngine,
		config.New,
		config.NewDBConfig,
		container.NewApp,
		database.New,
	)

	return nil, nil
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
