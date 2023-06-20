package elastic_client

import (
	"github.com/lowl11/lazy-elastic/internal/services/data_service"
	"github.com/lowl11/lazy-elastic/internal/services/index_service"
)

type Event struct {
	indexService *index_service.Service
	dataService  *data_service.Service
}

func Create(url string) *Event {
	return &Event{
		indexService: index_service.New(url),
		dataService:  data_service.New(url),
	}
}
