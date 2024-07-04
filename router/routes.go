package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaovit0r0/goOpportunities/handler"
)

// This function is responsible to creating and grouping
// the application route
func initializeRoutes(router *gin.Engine) {
	// Initialize
	handler.InitializeHandler()
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHadnler)
		v1.GET("/openings", handler.ListOpeningHadnler)
		v1.POST("/opening", handler.CreateOpeningHadnler)
		v1.PUT("/opening", handler.UpdateOpeningHadnler)
		v1.DELETE("/opening", handler.DeleteOpeningHadnler)
	}
}
