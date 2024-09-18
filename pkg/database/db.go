package database

import (
	"database/sql"
	"fmt"
	"inbody-ocr-backend/internal/app/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
)

type DB struct {
	*bun.DB
}

func New(cfg *config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	sqldb, err := sql.Open("nrmysql", dsn)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	return &DB{db}, nil
}
