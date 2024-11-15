package container

import (
	"errors"
	"fmt"
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/internal/controller"
	"inbody-ocr-backend/internal/infra/logging"
	"inbody-ocr-backend/internal/middleware"
	"inbody-ocr-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type container struct {
	userCtrl            *controller.UserController
	organizationCtrl    *controller.OrganizationController
	imageCtrl           *controller.ImageController
	imageDataCtrl       *controller.ImageDataController
	measurementDateCtrl *controller.MeasurementDateController
}

func NewCtrl(
	userCtrl *controller.UserController,
	organizationCtrl *controller.OrganizationController,
	imageCtrl *controller.ImageController,
	imageDataCtrl *controller.ImageDataController,
	measurementDateCtrl *controller.MeasurementDateController,
) *container {
	return &container{
		userCtrl:            userCtrl,
		organizationCtrl:    organizationCtrl,
		imageCtrl:           imageCtrl,
		imageDataCtrl:       imageDataCtrl,
		measurementDateCtrl: measurementDateCtrl,
	}
}

type App struct {
	r          *gin.Engine
	cfg        *config.Config
	db         *database.DB
	middleware *middleware.Middleware
}

func NewApp(r *gin.Engine, container *container, cfg *config.Config, db *database.DB, middleware *middleware.Middleware) *App {
	logging.Init()

	controller.SetUpRoutes(
		r,
		container.userCtrl,
		container.organizationCtrl,
		container.imageCtrl,
		container.imageDataCtrl,
		container.measurementDateCtrl,
		middleware,
	)

	return &App{
		r:          r,
		cfg:        cfg,
		db:         db,
		middleware: middleware,
	}
}

func (a *App) Migrate() error {
	return a.db.Migrate()
}

func (a *App) Seed() error {
	return a.db.SeedData()
}

func (a *App) Run() error {
	return a.r.Run(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
