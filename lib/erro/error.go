package erro

import "fmt"

type HttpError struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Url string `json:"url"`
	StatusCode int `json:"-"`
}

func NewHttpError(code int,statusCode int,msg string) *HttpError  {
	return &HttpError{
		Code: code,
		Msg: msg,
		StatusCode: statusCode,
	}
}

func (h HttpError) Error() string {
	switch m := h.Msg.(type) {
	case string:
		return m
	case map[string]string:
		total := ""
		for k, v := range m {
			total += fmt.Sprintf("%s: %s", k, v)
		}
		return total
	default:
		return ""
	}
}

func (h *HttpError) SetUrl(url string) *HttpError{
	h.Url = url
	return h
}

func (h *HttpError) SetMsg(msg interface{}) *HttpError{
	h.Msg = msg
	return h
}

func (h *HttpError) SetHttpCode(statusCode int) *HttpError {
	h.StatusCode = statusCode
	return h
}
