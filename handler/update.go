package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Oppening",
	})
}
