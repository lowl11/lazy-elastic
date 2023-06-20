package es_api

import (
	"github.com/lowl11/lazy-elastic/elastic_client"
	"github.com/lowl11/lazy-elastic/elastic_search"
)

func NewClient(baseURL string) *elastic_client.Event {
	return elastic_client.Create(baseURL)
}

func NewSearch[T any](baseURL string) *elastic_search.Event[T] {
	return elastic_search.Create[T](baseURL)
}

//func BulkArray[T any](bulkArray []T) []any {
//	return type_list.NewWithList[T, any](bulkArray...).Select(func(item T) any {
//		return item
//	}).Slice()
//}
