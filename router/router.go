package router

import "github.com/gin-gonic/gin"

// Qualquer função ou varíavel que deva ser visível em outro arquivo
// deve ser criada com letra maiúscula, OBRIGATORIAMENTE
func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
