package search_service

import (
	"bytes"
	"encoding/json"
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/requests"
	"net/http"
)

func (service *Service[T]) All() *Service[T] {
	service.body = map[string]any{
		"query": map[string]any{
			"match_all": make(map[string]any),
		},
	}

	return service
}

func (service *Service[T]) MultiMatch(query string, fields []string) *Service[T] {
	service.body = map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"fields": fields,
				"query":  query,
			},
		},
	}

	return service
}

func (service *Service[T]) Search() ([]T, error) {
	queryInBytes, err := json.Marshal(service.body)
	if err != nil {
		return nil, err
	}

	response, err := requests.New(
		http.MethodPost,
		service.baseURl+"/"+service.indexName+"/_search",
		bytes.NewBuffer(queryInBytes)).
		Header("Content-Type", "application/json").
		Send()
	if err != nil {
		return nil, err
	}

	result := es_model.SearchResponse[T]{}
	if err = json.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return type_list.NewWithList[es_model.SearchHit[T], T](result.Hits.Hits...).Select(
		func(item es_model.SearchHit[T]) T {
			return item.Source
		}).Slice(), nil
}
