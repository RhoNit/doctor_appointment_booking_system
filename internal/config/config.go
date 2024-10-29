package config

import "os"

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type CacheConfig struct{}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
