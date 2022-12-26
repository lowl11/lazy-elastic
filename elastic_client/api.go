package elastic_client

import (
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/data_service"
	"github.com/lowl11/lazy-elastic/es_services/index_service"
)

func (event *Event) AllIndices() ([]string, error) {
	indices, err := index_service.All(event.baseURL + "/_cat/indices?format=json")
	if err != nil {
		return nil, err
	}

	return type_list.NewWithList[es_model.IndexGet, string](indices...).
		Select(func(item es_model.IndexGet) string {
			return item.Name
		}).Slice(), nil
}

func (event *Event) CreateIndex(indexName string, mappings map[string]any) error {
	if index_service.Exist(event.baseURL, indexName) {
		return nil
	}

	return index_service.Create(event.baseURL, indexName, mappings)
}

func (event *Event) DeleteIndex(indexName string) error {
	return index_service.Delete(event.baseURL, indexName)
}

func (event *Event) Insert(id, indexName string, object any) error {
	return data_service.Insert(id, object, event.baseURL, indexName)
}

func (event *Event) InsertMultiple(indexName string, objects []any) error {
	return data_service.InsertMultiple(event.baseURL, indexName, objects)
}

func (event *Event) DeleteItem(indexName, id string) error {
	return data_service.Delete(event.baseURL, indexName, id)
}
