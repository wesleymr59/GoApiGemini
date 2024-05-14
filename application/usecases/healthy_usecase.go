package usecases

type HealthyUsecase struct{}

func NewHealthyService() *HealthyUsecase {
	return &HealthyUsecase{}
}

func (s *HealthyUsecase) GetHelloMessage() string {
	return "Hello, World!"
}
