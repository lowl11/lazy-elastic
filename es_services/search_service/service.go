package search_service

type Service[T any] struct {
	baseURl   string
	indexName string

	body map[string]any
}

func New[T any](baseURL, indexName string) *Service[T] {
	return &Service[T]{
		baseURl:   baseURL,
		indexName: indexName,
	}
}
