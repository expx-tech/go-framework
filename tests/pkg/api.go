package types

import (
	"net/http"

	lib "github.com/expx-tech/go-framework/tests/pkg/internal"
)

type Api struct {
	Schema  string
	Host    string
	Headers map[string]string
}

func NewApi() Api {
	return Api{
		Schema: "http",
		Host:   "127.0.0.1:8080",
	}
}

func (a Api) GetHttp(url string, result *http.Response) error {
	return lib.GetHttp(a.Schema+"://"+a.Host+url, result)
}

func (a Api) PostHttp(url string, contentType string, body []byte, result *http.Response) error {
	return lib.PostHttp(a.Schema+"://"+a.Host+url, contentType, body, a.Headers, result)
}
