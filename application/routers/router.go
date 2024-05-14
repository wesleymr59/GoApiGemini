package routers

import (
	"api-gemini-go/application/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	helloController := interfaces.NewHelloHandler()
	router.GET("/hello", helloController.HandleHello)

	return router
}
