package search_service

type Service[T any] struct {
	baseURl   string
	indexName string

	prefixLength  int
	maxExpansions int
	fuzziness     string
	size          int

	body map[string]any

	isMultiMatch bool
}

func New[T any](baseURL, indexName string) *Service[T] {
	return &Service[T]{
		baseURl:   baseURL,
		indexName: indexName,

		prefixLength:  defaultPrefixLength,
		maxExpansions: defaultMaxExpansions,
		fuzziness:     defaultFuzziness,
	}
}
