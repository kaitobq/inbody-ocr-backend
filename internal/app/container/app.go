package container

import (
	"errors"
	"fmt"
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/internal/controller"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type container struct {
	userCtrl *controller.UserController
	organizationCtrl *controller.OrganizationController
	membershipCtrl *controller.UserOrganizationMembershipController
	imageCtrl *controller.ImageController
	tokenService service.TokenService
}

func NewCtrl(
	userCtrl *controller.UserController,
	organizationCtrl *controller.OrganizationController,
	membershipCtrl *controller.UserOrganizationMembershipController,
	imageCtrl *controller.ImageController,
	tokenService service.TokenService,
) *container {
	return &container{
		userCtrl: userCtrl,
		organizationCtrl: organizationCtrl,
		membershipCtrl: membershipCtrl,
		imageCtrl: imageCtrl,
		tokenService: tokenService,
	}
}

type App struct {
	r *gin.Engine
	cfg *config.Config
	db     *database.DB
}

func NewApp(r *gin.Engine, container *container, cfg *config.Config, db *database.DB) *App {
	controller.SetUpRoutes(
		r,
		container.userCtrl,
		container.organizationCtrl,
		container.membershipCtrl,
		container.imageCtrl,
		container.tokenService,
	)

	return &App{
		r: r,
		cfg: cfg,
		db: db,
	}
}

func (a *App) Migrate() error {
	return a.db.Migrate()
}

func (a *App) Run() error {
	return a.r.Run(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
