package jsonschema

var schemaMapper = map[string]interface{}{
	"hatena_hotentry/:date":      new(HatenaHotEntryAPISchema),
	"hatena_hotentry_detail/:id": new(HatenaHotEntryDetailAPISchema),
	"qiita_entry/:date":          new(QiitaEntryAPISchema),
	"qiita_entry_detail/:id":     new(QiitaEntryDetailAPISchema),
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
