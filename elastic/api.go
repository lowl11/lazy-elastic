package elastic

import (
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/index_service"
)

func (event *Event) AllIndices() ([]string, error) {
	indices, err := index_service.AllIndices(event.baseURL + "/_cat/indices?format=json")
	if err != nil {
		return nil, err
	}

	return type_list.NewWithList[es_model.IndexGet, string](indices...).
		Select(func(item es_model.IndexGet) string {
			return item.Name
		}).Slice(), nil
}
