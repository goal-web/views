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
	app.Singleton("view", func(config contracts.Config) contracts.Views {
		conf, _ := config.Get("views").(Config)
		return NewView(conf.Path)
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {

}
