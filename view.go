package views

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"path/filepath"
	"strings"
	"sync"
)

type View struct {
	templates map[string]*pongo2.Template
	debug     bool

	path string

	mutex sync.Mutex
}

func (v *View) Register(name, template string) {
	tpl, err := pongo2.FromString(template)

	if err != nil {
		logs.Default().
			WithField("name", name).
			WithError(err).
			WithField("template", template).
			Debug(fmt.Sprintf("Failed to parse template: %s", err))

		panic(RegisterException{
			Exception: exceptions.WithError(err),
			Name:      name,
			Template:  template,
		})
	}

	v.mutex.Lock()
	v.templates[name] = tpl
	v.mutex.Unlock()
}

func NewView(path string, debug bool) contracts.Views {
	return &View{
		templates: make(map[string]*pongo2.Template),
		path:      path,
		debug:     debug,
	}
}

func (v *View) Render(name string, data ...any) contracts.HttpResponse {
	context, err := utils.ToFields(utils.DefaultValue(data, any(contracts.Fields{})))
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

	tpl, exists := v.templates[name]
	if !exists {
		viewPath := name
		if !strings.HasPrefix(name, "/") {
			viewPath = filepath.Join(v.path, name)
		}

		tpl, err = pongo2.FromFile(viewPath)

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

		if !v.debug {
			v.mutex.Lock()
			v.templates[name] = tpl
			v.mutex.Unlock()
		}
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
	return NewResponse([]byte(output))
}
