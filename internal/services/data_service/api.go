package data_service

import (
	"encoding/json"
	"errors"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/internal/services/requests"
	"net/http"
)

func (service *Service) Insert(id string, object any, indexName string) error {
	if object == nil {
		return errors.New("object is null")
	}

	response, statusCode, err := requests.
		New(http.MethodPost, service.url+"/"+indexName+"/_doc/"+id, object).
		Header("Content-Type", "application/json").
		SendStatus()
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		errorObject := es_model.Error{}
		if err = json.Unmarshal(response, &errorObject); err != nil {
			return errors.New(string(response))
		}
		return errors.New(errorObject.Error.Reason)
	}

	return nil
}

func (service *Service) InsertMultiple(indexName string, objects []es_model.InsertMultipleData) error {
	if objects == nil {
		return errors.New("object is null")
	}

	if len(objects) == 0 {
		return nil
	}

	var bulkObjects string
	for _, obj := range objects {
		insertModel := &es_model.InsertData{
			Index: struct {
				ID   string `json:"_id"`
				Name string `json:"_index"`
				Type string `json:"_type"`
			}{ID: obj.ID, Name: indexName, Type: "_doc"},
		}

		insertObjectInBytes, err := json.Marshal(insertModel)
		if err != nil {
			return err
		}

		objectInBytes, err := json.Marshal(obj.Object)
		if err != nil {
			return err
		}

		bulkObjects += string(insertObjectInBytes) + "\n"
		bulkObjects += string(objectInBytes) + "\n"
	}

	response, status, err := requests.
		New(http.MethodPost, service.url+"/_bulk", bulkObjects).
		Header("Content-Type", "application/x-ndjson").
		SendStatus()
	if err != nil {
		return err
	}

	if status != http.StatusOK && status != http.StatusCreated {
		errorObject := es_model.Error{}
		if err = json.Unmarshal(response, &errorObject); err != nil {
			return errors.New(string(response))
		}
		return errors.New(errorObject.Error.Reason)
	}

	return nil
}

func (service *Service) Delete(indexName, id string) error {
	response, status, err := requests.
		New(http.MethodDelete, service.url+"/"+indexName+"/_doc/"+id, nil).
		SendStatus()
	if err != nil {
		return err
	}

	if status != http.StatusOK && status != http.StatusCreated {
		errorObject := es_model.Error{}
		if err = json.Unmarshal(response, &errorObject); err != nil {
			return errors.New(string(response))
		}
		return errors.New(errorObject.Error.Reason)
	}

	return nil
}
