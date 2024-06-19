package router

import "github.com/gin-gonic/gin"

// Qualquer função ou varíavel que deva ser visível em outro arquivo
// deve ser criada com letra maiúscula, OBRIGATORIAMENTE
func Initialize() {
	// Inicializando o Router
	router := gin.Default()

	// Inicializando as Rotas
	initializeRoutes(router)

	// Inicializando o servidor
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
