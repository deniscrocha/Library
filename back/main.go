package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getInit(context *gin.Context) {
	greeting := "Olá, mundo!"
	context.JSON(http.StatusOK, gin.H{"message": *&greeting})
}

func main() {
	router := gin.Default()
	router.GET("/", getInit)
	router.Run("localhost:8080")
}
