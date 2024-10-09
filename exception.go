package views

import "github.com/goal-web/contracts"

type ViewNotFoundException struct {
	contracts.Exception

	Name string
	Data any
}

type DataInvalidException struct {
	contracts.Exception

	Name string
	Data any
}

type ViewRenderException struct {
	contracts.Exception

	Name string
	Data any
}

type RegisterException struct {
	contracts.Exception

	Name     string
	Template string
}
