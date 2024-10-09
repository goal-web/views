package views

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
)

func NewResponse(contents []byte) contracts.HttpResponse {
	return &Response{
		BaseResponse: http.NewBaseResponse(200, map[string][]string{
			"Content-Type": {"text/html; charset=utf-8"},
		}),
		bytes: contents,
	}
}

type Response struct {
	*http.BaseResponse

	bytes []byte
}

func (res *Response) Bytes() []byte {
	return res.bytes
}
