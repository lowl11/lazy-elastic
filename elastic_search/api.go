package elastic_search

import "github.com/lowl11/lazy-elastic/es_services/search_service"

func (event *Event[T]) All(indexName string) ([]T, error) {
	results, err := search_service.New[T](event.baseURL, indexName).
		All().
		Search()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (event *Event[T]) MultiMatch(indexName, query string, fields []string) ([]T, error) {
	results, err := search_service.New[T](event.baseURL, indexName).
		MultiMatch(query, fields).
		Search()
	if err != nil {
		return nil, err
	}

	return results, nil
}
