package requests

import "net/http"

func (service *Service) fillCookies() {
	for key, value := range service.cookies {
		service.request.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
}

func (service *Service) fillHeaders() {
	for key, value := range service.headers {
		service.request.Header.Set(key, value)
	}
}

func (service *Service) fillBasicAuth() {
	if !service.isBasicAuth {
		return
	}

	service.request.SetBasicAuth(service.username, service.password)
}

func (service *Service) httpClient() *http.Client {
	return &http.Client{}
}
