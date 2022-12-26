package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var buffer *bytes.Buffer
	if service.body != nil {
		if stringValue, isString := service.body.(string); isString {
			buffer = bytes.NewBuffer([]byte(stringValue))
		} else {
			bodyInBytes, err := json.Marshal(service.body)
			if err != nil {
				return nil, err
			}

			buffer = bytes.NewBuffer(bodyInBytes)
		}
	}

	// create request
	var request *http.Request
	var err error
	if service.body != nil {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, buffer)
	} else {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, nil)
	}
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
