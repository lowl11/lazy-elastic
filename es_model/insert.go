package es_model

type InsertData struct {
	Index struct {
		Name string `json:"_index"`
		Type string `json:"_type"`
	} `json:"index"`
}
