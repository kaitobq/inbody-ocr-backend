package container

import (
	"errors"
	"fmt"
	"inbody-ocr-backend/internal/app/config"
	"inbody-ocr-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type App struct {
	r *gin.Engine
	cfg *config.Config
	db     *database.DB
}

func NewApp(r *gin.Engine, cfg *config.Config, db *database.DB) *App {
	return &App{
		r: r,
		cfg: cfg,
		db: db,
	}
}

func (a *App) Run() error {
	return a.r.Run(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
