package requests

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (service *Service) Header(key, value string) *Service {
	service.headers[key] = value
	return service
}

func (service *Service) Headers(headers map[string]string) *Service {
	service.headers = headers
	return service
}

func (service *Service) Cookie(key, value string) *Service {
	service.cookies[key] = value
	return service
}

func (service *Service) Cookies(cookies map[string]string) *Service {
	service.cookies = cookies
	return service
}

func (service *Service) Basic(username, password string) *Service {
	service.isBasicAuth = true
	service.username = username
	service.password = password
	return service
}

func (service *Service) WithLogs() {
	service.withLogs = true
}

func (service *Service) Send() ([]byte, error) {
	request, err := http.NewRequest(service.method, service.url, nil)
	if err != nil {
		return nil, err
	}

	service.request = request

	service.fillCookies()
	service.fillHeaders()
	service.fillBasicAuth()

	client := service.httpClient()

	response, err := client.Do(service.request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			if service.withLogs {
				log.Println("Close Elasticsearch request error: ", err)
			}
		}
	}()

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseInBytes, nil
}
