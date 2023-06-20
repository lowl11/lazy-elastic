package elastic_client

import (
	"github.com/lowl11/lazy-elastic/es_model"
)

func (event *Event) AllIndices() ([]string, error) {
	indices, err := event.indexService.All()
	if err != nil {
		return nil, err
	}

	indexNames := make([]string, 0, len(indices))
	for _, index := range indices {
		if index.Name[0] == '.' {
			continue
		}

		indexNames = append(indexNames, index.Name)
	}

	return indexNames, nil
}

func (event *Event) CreateIndex(indexName string, mappings map[string]any) error {
	if event.indexService.Exist(indexName) {
		return nil
	}

	return event.indexService.Create(indexName, mappings)
}

func (event *Event) ExistIndex(indexName string) bool {
	return event.indexService.Exist(indexName)
}

func (event *Event) DeleteIndex(indexName string) error {
	return event.indexService.Delete(indexName)
}

func (event *Event) Insert(id, indexName string, object any) error {
	return event.dataService.Insert(id, object, indexName)
}

func (event *Event) InsertMultiple(indexName string, objects []es_model.InsertMultipleData) error {
	return event.dataService.InsertMultiple(indexName, objects)
}

func (event *Event) DeleteItem(indexName, id string) error {
	return event.dataService.Delete(indexName, id)
}
