package index_service

import (
	"encoding/json"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/internal/services/requests"
	"net/http"
)

func (service *Service) All() ([]es_model.IndexGet, error) {
	response, err := requests.New(http.MethodGet, service.url+"/_cat/indices?format=json", nil).
		Send()
	if err != nil {
		return nil, err
	}

	list := make([]es_model.IndexGet, 0)
	if err = json.Unmarshal(response, &list); err != nil {
		return nil, err
	}

	return list, nil
}

func (service *Service) Exist(indexName string) bool {
	_, status, err := requests.New(http.MethodGet, service.url+"/"+indexName, nil).
		SendStatus()
	if err != nil {
		return false
	}

	return status != http.StatusNotFound
}

func (service *Service) Create(indexName string, mappings map[string]any) error {
	_, err := requests.New(http.MethodPut, service.url+"/"+indexName, nil).
		Send()
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) Delete(indexName string) error {
	_, err := requests.New(http.MethodDelete, service.url+"/"+indexName, nil).
		Send()
	if err != nil {
		return err
	}

	return nil
}
