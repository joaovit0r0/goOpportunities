package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Oppening",
	})
}
