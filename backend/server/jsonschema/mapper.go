package jsonschema

type APISchema interface {
	getRequestSchema() map[string]interface{}
	getResponseSchema() map[string]interface{}
}

var schemaMapper = map[string]interface{}{
	"hatena_hotentry_detail/:id": new(HatenaHotEntryDetailAPISchema),
}

func getSchemaMapper() map[string]interface{} {
	allSchemaMapper := make(map[string]interface{})
	allSchemaList := []map[string]interface{}{schemaMapper}

	for _, schem := range allSchemaList {
		for k, v := range schem {
			allSchemaMapper[k] = v
		}
	}
	return allSchemaMapper
}
