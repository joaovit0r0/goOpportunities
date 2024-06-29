package config

import (
	"os"

	"github.com/joaovit0r0/goOpportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// 	Check if database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		// create the database directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		// create database file
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.ErrorF("db open connection error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{}) // cria
	if err != nil {
		logger.ErrorF("db automigrate erro: %v", err)
		return nil, err
	}
	return db, nil
}
