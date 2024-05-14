package interfaces

import (
	"api-gemini-go/application/composers"

	"github.com/gin-gonic/gin"
)

type HealthyHandler struct {
	composer composers.HealthyComposer
}

func NewHelloHandler() *HealthyHandler {
	return &HealthyHandler{
		composer: *composers.NewHealthyComposer(),
	}
}

func (h *HealthyHandler) HandleHello(c *gin.Context) {
	service := h.composer.ComposeHelloService()
	message := service.GetHelloMessage()
	c.JSON(200, gin.H{"message": message})
}
