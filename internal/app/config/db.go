package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     load("DB_HOST"),
		Port:     load("DB_PORT"),
		User:	  load("DB_USER"),
		Password: load("DB_PASSWORD"),
		DBName:   load("DB_NAME"),
	}
}

func load(key string) string {
	return os.Getenv(key)
}
