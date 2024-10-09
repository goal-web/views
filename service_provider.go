package views

import (
	"github.com/goal-web/contracts"
)

func NewService() contracts.ServiceProvider {
	return &ServiceProvider{}
}

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("view", func() contracts.Views {
		return NewView()
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {

}
