package views

import "github.com/goal-web/http"

type Response struct {
	http.BaseResponse

	bytes []byte
}

func (res *Response) Bytes() []byte {
	return res.bytes
}
