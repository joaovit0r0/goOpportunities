package main

import (
	"github.com/joaovit0r0/goOpportunities/config"
	"github.com/joaovit0r0/goOpportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("config inizilation error: %v", err)
		return
	}

	router.Initialize()
}
