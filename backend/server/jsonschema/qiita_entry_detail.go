package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
)

type QiitaEntryDetailAPISchema struct{}

func (s QiitaEntryDetailAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Qiita Entry Request API ",
		"description": "Qiita Entry Request description",
		"method":      "GET",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]string{
				"type":        "integer",
				"example":     "1",
				"description": "Qiita Entry ID",
			},
		},
		"required": []string{
			"id",
		},
		"additionalProperties": false,
	}
	return schema
}

func (s QiitaEntryDetailAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Qiita Entry Response API",
		"description": "Qiita Entry Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"qiita_entry": qiitaentry.Schema(),
		},
		"required": []string{
			"qiita_entry",
		},
		"additionalProperties": false,
	}
	return schema
}
