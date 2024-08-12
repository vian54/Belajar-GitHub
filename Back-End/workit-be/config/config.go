package config

import (
	"gorm.io/gorm"
)

type Config struct {
	App App
	DB  *gorm.DB
}

func CreateNewConfig() *Config {
	return &Config{}
}
