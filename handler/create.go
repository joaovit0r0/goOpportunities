package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
