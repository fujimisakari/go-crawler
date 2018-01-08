package jsonschema

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
