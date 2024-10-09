package views

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"sync"
)

type View struct {
	templates map[string]*pongo2.Template

	mutex sync.Mutex
}

func NewView() contracts.Views {
	return &View{
		templates: make(map[string]*pongo2.Template),
	}
}

func (v *View) Render(name string, data any) contracts.HttpResponse {
	var err error
	tpl, exists := v.templates[name]
	if !exists {
		tpl, err = pongo2.FromFile(name)
		if err != nil {
			logs.Default().
				WithField("name", name).
				WithField("data", data).
				Debug(fmt.Sprintf("template %s not found", name))
			panic(ViewNotFoundException{
				Exception: exceptions.WithError(err),
				Name:      name,
				Data:      data,
			})
		}
		v.mutex.Lock()
		v.templates[name] = tpl
		v.mutex.Unlock()
	}

	context, err := utils.ToFields(data)

	if err != nil {
		logs.Default().
			WithField("name", name).
			WithField("data", data).
			Debug(fmt.Sprintf("Failed to parse data: %s", err.Error()))
		panic(DataInvalidException{
			Exception: exceptions.WithError(err),
			Name:      name,
			Data:      data,
		})
	}

	output, err := tpl.Execute(pongo2.Context(context))
	if err != nil {
		logs.Default().
			WithField("name", name).
			WithField("data", data).
			Debug(fmt.Sprintf("Failed to render view: %s", err.Error()))
		panic(ViewRenderException{
			Exception: exceptions.WithError(err),
			Name:      name,
			Data:      data,
		})
	}
	return &Response{bytes: []byte(output)}
}
