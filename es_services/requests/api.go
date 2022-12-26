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
	// create request
	request, err := http.NewRequest(service.method, service.url, nil)
	if err != nil {
		return nil, err
	}

	service.request = request

	// fill need data
	service.fillCookies()
	service.fillHeaders()
	service.fillBasicAuth()

	// get http client
	client := service.httpClient()

	// send request
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

	// write status
	service.status = response.StatusCode

	// read response to bytes
	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		service.err = err
		return nil, err
	}

	// write response
	service.response = responseInBytes

	return responseInBytes, nil
}

func (service *Service) SendWithStatus() ([]byte, int, error) {
	_, _ = service.Send()
	return service.response, service.status, service.err
}
