package search_service

import (
	"encoding/json"
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/requests"
	"net/http"
	"strconv"
)

func (service *Service[T]) Prefix(length int) *Service[T] {
	if length > 2 {
		service.prefixLength = length
	}

	return service
}

func (service *Service[T]) MaxExpansions(maxExpansions int) *Service[T] {
	if maxExpansions > 1 {
		service.maxExpansions = maxExpansions
	}

	return service
}

func (service *Service[T]) Fuzziness(fuzziness int) *Service[T] {
	if fuzziness > 2 {
		service.fuzziness = strconv.Itoa(fuzziness)
	}

	return service
}

func (service *Service[T]) All() *Service[T] {
	service.body = map[string]any{
		"query": map[string]any{
			"match_all": make(map[string]any),
		},
	}

	return service
}

func (service *Service[T]) Not(conditions ...map[string]any) *Service[T] {
	//service.body["query"].(map[string]any)["bool"].(map[string]any)["must_not"] = conditions
	service.body["query"].(map[string]any)["bool"].(map[string]any)["must_not"] = []map[string]any{
		{
			"bool": map[string]any{
				"filter": conditions,
			},
		},
	}
	return service
}

func (service *Service[T]) MultiMatch(query string, fields []string) *Service[T] {
	service.isMultiMatch = true

	service.body = map[string]any{
		"query": map[string]any{
			"bool": map[string]any{
				"must": []map[string]any{
					{
						"multi_match": map[string]any{
							"fields": fields,
							"query":  query,
						},
					},
				},
			},
		},
	}

	return service
}

func (service *Service[T]) Search() ([]T, error) {
	service.fillAttributes()

	response, err := requests.New(
		http.MethodPost,
		service.baseURl+"/"+service.indexName+"/_search",
		service.body,
	).
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
