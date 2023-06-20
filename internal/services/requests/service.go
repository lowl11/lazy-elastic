package requests

import (
	"net/http"
	"time"
)

type Service struct {
	method string
	url    string
	body   any

	headers map[string][]string
	cookies map[string]string

	isXml bool

	timeout       time.Duration
	customTimeout bool

	isBasicAuth bool
	username    string
	password    string

	request *http.Request

	sslCheck bool
	noProxy  bool

	status int
}

func New(method, url string, body any) *Service {
	return &Service{
		method: method,
		url:    url,
		body:   body,

		headers: make(map[string][]string),
		cookies: make(map[string]string),
	}
}

func NewSoap(url string, body any) *Service {
	return &Service{
		method: http.MethodPost,
		url:    url,
		body:   body,

		headers: make(map[string][]string),
		cookies: make(map[string]string),
	}
}
