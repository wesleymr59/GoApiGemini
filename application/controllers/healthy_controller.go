package controllers

import (
	"api-gemini-go/application/composers"

	"github.com/gin-gonic/gin"
)

type HealthyController struct {
	composer composers.HealthyComposer
}

func NewHealthyController() *HealthyController {
	return &HealthyController{
		composer: *composers.NewHealthyComposer(),
	}
}

func (c *HealthyController) HandleHealty(ctx *gin.Context) {
	service := c.composer.ComposeHelloService()
	message := service.GetHelloMessage()

	ctx.JSON(200, gin.H{"message": message})
}
