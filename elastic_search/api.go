package elastic_search

import "github.com/lowl11/lazy-elastic/es_services/search_service"

func (event *Event[T]) All(indexName string) *search_service.Service[T] {
	return search_service.New[T](event.baseURL, indexName).All()
}

func (event *Event[T]) MultiMatch(indexName, query string, fields []string) *search_service.Service[T] {
	return search_service.New[T](event.baseURL, indexName).MultiMatch(query, fields)
}
