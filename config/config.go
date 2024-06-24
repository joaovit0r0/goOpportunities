package config

import (
	"gorm.io/gorm"
)

// Declaração uma variável global
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
