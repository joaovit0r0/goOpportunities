package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Oppening",
	})
}

func ListOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Oppenings",
	})
}

func CreateOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Oppening",
	})
}

func UpdateOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Oppening",
	})
}

func DeleteOppeningHadnler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Oppening",
	})
}
