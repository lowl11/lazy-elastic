package index_service

import (
	"encoding/json"
	"github.com/lowl11/lazy-elastic/es_model"
	"github.com/lowl11/lazy-elastic/es_services/requests"
	"net/http"
)

func AllIndices(url string) ([]es_model.IndexGet, error) {
	response, err := requests.New(http.MethodGet, url, nil).Send()
	if err != nil {
		return nil, err
	}

	list := make([]es_model.IndexGet, 0)
	if err = json.Unmarshal(response, &list); err != nil {
		return nil, err
	}

	return list, nil
}
