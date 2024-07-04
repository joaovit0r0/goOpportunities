package handler

import (
	"github.com/joaovit0r0/goOpportunities/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
