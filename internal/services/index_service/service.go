package index_service

type Service struct {
	url string
}

func New(url string) *Service {
	return &Service{
		url: url,
	}
}
