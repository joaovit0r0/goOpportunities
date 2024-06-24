package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaovit0r0/goOpportunities/handler"
)

/*
Função responsável por criar e agrupar as rotas da aplicação
*/
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHadnler)
		v1.GET("/openings", handler.ListOpeningHadnler)
		v1.POST("/opening", handler.CreateOpeningHadnler)
		v1.PUT("/opening", handler.UpdateOpeningHadnler)
		v1.DELETE("/opening", handler.DeleteOpeningHadnler)
	}
}
