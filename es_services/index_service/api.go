package index_service

import (
	"encoding/json"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/requests"
	"net/http"
)

func All(url string) ([]es_model.IndexGet, error) {
	response, err := requests.New(http.MethodGet, url, nil).Send()
	if err != nil {
		return nil, err
	}

	list := make([]es_model.IndexGet, 0)
	if err = json.Unmarshal(response, &list); err != nil {
		return nil, err
	}

	return list, nil
}

func Exist(url, indexName string) bool {
	_, status, err := requests.New(http.MethodGet, url+"/"+indexName, nil).SendWithStatus()
	if err != nil {
		return false
	}

	return status != http.StatusNotFound
}

func Create(url, indexName string, mappings map[string]any) error {
	_, err := requests.New(http.MethodPut, url+"/"+indexName, nil).Send()
	if err != nil {
		return err
	}

	return nil
}

func Delete(url, indexName string) error {
	_, err := requests.New(http.MethodDelete, url+"/"+indexName, nil).Send()
	if err != nil {
		return err
	}

	return nil
}
