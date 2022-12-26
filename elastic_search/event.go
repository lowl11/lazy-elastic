package elastic_search

type Event[T any] struct {
	baseURL string
}

func Create[T any](baseURL string) *Event[T] {
	return &Event[T]{
		baseURL: baseURL,
	}
}
