package es_api

import "github.com/lowl11/lazy-elastic/elastic"

func New(baseURl string) *elastic.Event {
	return elastic.Create(baseURl)
}
