package user_service

type Application struct {
	Address string
}

func NewApplication(address string) *Application {
	return &Application{
		Address: address,
	}
}
