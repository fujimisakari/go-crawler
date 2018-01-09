package qiitaentry

func Schema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Qiita Entry item",
		"description": "Qiita Entry info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "integer",
				"example": "1",
			},
			"user": map[string]interface{}{
				"type":    "string",
				"example": "Qiita User",
			},
			"title": map[string]interface{}{
				"type":    "string",
				"example": "Qiita Entry Title",
			},
			"link": map[string]interface{}{
				"type":    "string",
				"example": "Url",
			},
			"posted_at": map[string]interface{}{
				"type":    "string",
				"example": "Qiita Posted At",
			},
		},
	}
	return schema
}
