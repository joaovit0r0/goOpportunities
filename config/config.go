package config

import (
	"fmt"

	"gorm.io/gorm"
)

// Declaração uma variável global
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initialazing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
