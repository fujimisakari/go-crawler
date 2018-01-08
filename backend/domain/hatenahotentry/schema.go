package hatenahotentry

func Schema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Hatena Hotentry item",
		"description": "Hatena Hotentry info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "integer",
				"example": "1",
			},
			"title": map[string]interface{}{
				"type":    "string",
				"example": "Hatena Hotentry Title",
			},
			"link": map[string]interface{}{
				"type":    "string",
				"example": "Url",
			},
			"description": map[string]interface{}{
				"type":    "string",
				"example": "Hatena Hotentry Description",
			},
		},
	}
	return schema
}
