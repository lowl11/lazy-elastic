package es_model

type InsertData struct {
	Index struct {
		ID   string `json:"_id"`
		Name string `json:"_index"`
		Type string `json:"_type"`
	} `json:"index"`
}

type InsertMultipleData struct {
	ID     string
	Object any
}
