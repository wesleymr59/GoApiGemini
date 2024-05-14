package composers

import (
	"api-gemini-go/application/usecases"
)

type HealthyComposer struct{}

func NewHealthyComposer() *HealthyComposer {
	return &HealthyComposer{}
}

func (c *HealthyComposer) ComposeHelloService() *usecases.HealthyUsecase {
	return usecases.NewHealthyService()
}
