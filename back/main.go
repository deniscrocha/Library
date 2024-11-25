package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.com/api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func getInit(context *gin.Context) {
	greeting := "Olá, mundo!"
	context.JSON(http.StatusOK, gin.H{"message": *&greeting})
}

func main() {
	router := gin.Default()
	router.GET("/", getInit)
	router.Run()
}
