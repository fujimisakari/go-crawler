package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
)

type HatenaHotEntryAPISchema struct{}

func (s HatenaHotEntryAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Hatena Hotentry Request API ",
		"description": "Hatena Hotentry Request description",
		"method":      "GET",
		"type":        "object",
		"properties": map[string]interface{}{
			"date": map[string]string{
				"type":        "string",
				"example":     "2018-1-1",
				"description": "crawl date",
			},
		},
		"required": []string{
			"date",
		},
		"additionalProperties": false,
	}
	return schema
}

func (s HatenaHotEntryAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Hatena Hotentry Response API",
		"description": "Hatena Hotentry Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"hatena_hotentries": map[string]interface{}{
				"type":  "array",
				"items": hatenahotentry.Schema(),
			},
		},
		"required": []string{
			"hatena_hotentries",
		},
		"additionalProperties": false,
	}
	return schema
}
