package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Função responsável por criar e agrupar as rotas da aplicação
*/
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Oppening",
			})
		})
		v1.GET("/oppenings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Oppenings",
			})
		})
		v1.POST("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Oppening",
			})
		})
		v1.PUT("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Oppening",
			})
		})
		v1.DELETE("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Oppening",
			})
		})
	}
}
