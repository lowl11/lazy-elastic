package elastic_client

type Event struct {
	baseURL string
}

func Create(baseURL string) *Event {
	return &Event{
		baseURL: baseURL,
	}
}
