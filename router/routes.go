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
		v1.GET("/oppening", handler.ShowOppeningHadnler)
		v1.GET("/oppenings", handler.ListOppeningHadnler)
		v1.POST("/oppening", handler.CreateOppeningHadnler)
		v1.PUT("/oppening", handler.UpdateOppeningHadnler)
		v1.DELETE("/oppening", handler.DeleteOppeningHadnler)
	}
}
