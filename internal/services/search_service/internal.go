package search_service

const (
	defaultPrefixLength  = 2
	defaultMaxExpansions = 1
	defaultFuzziness     = "AUTO"
)

func (service *Service[T]) fillAttributes() {
	if service.isMultiMatch {
		service.fillMultiMatch()
	}
}

func (service *Service[T]) fillMultiMatch() {
	// multi match configs
	multiMatch := service.body["query"].(map[string]any)["bool"].(map[string]any)["must"].([]map[string]any)[0]["multi_match"].(map[string]any)

	multiMatch["prefix_length"] = service.prefixLength
	multiMatch["max_expansions"] = service.maxExpansions
	multiMatch["fuzziness"] = service.fuzziness
}
