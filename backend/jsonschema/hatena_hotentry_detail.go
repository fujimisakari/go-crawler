package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
)

type HatenaHotEntryDetailAPISchema struct{}

func (s HatenaHotEntryDetailAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Hatena Hotentry Request API ",
		"description": "Hatena Hotentry Request description",
		"method":      "GET",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]string{
				"type":        "integer",
				"example":     "1",
				"description": "Hatena Hotentry ID",
			},
		},
		"required": []string{
			"id",
		},
		"additionalProperties": false,
	}
	return schema
}

func (s HatenaHotEntryDetailAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Hatena Hotentry Response API",
		"description": "Hatena Hotentry Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"hatena_hotentry": hatenahotentry.Schema(),
		},
		"required": []string{
			"hatena_hotentry",
		},
		"additionalProperties": false,
	}
	return schema
}
