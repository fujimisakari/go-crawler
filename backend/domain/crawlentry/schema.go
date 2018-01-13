package crawlentry

func Schema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Crawl Entry item",
		"description": "Crawl Entry info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "integer",
				"example": "1",
			},
			"date": map[string]interface{}{
				"type":    "string",
				"example": "Crawl Date",
			},
			"crawl_items": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type":    "string",
					"example": "hatena_hotentry",
				},
			},
		},
	}
	return schema
}
