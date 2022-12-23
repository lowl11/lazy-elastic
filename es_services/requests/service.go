package requests

import "net/http"

type Service struct {
	method string
	url    string
	body   any

	cookies map[string]string
	headers map[string]string

	isBasicAuth bool
	username    string
	password    string

	withLogs bool

	request *http.Request
}

func New(method, url string, body any) *Service {
	return &Service{
		method: method,
		url:    url,
		body:   body,

		cookies: make(map[string]string),
		headers: make(map[string]string),
	}
}
