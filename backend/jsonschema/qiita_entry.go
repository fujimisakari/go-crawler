package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
)

type QiitaEntryAPISchema struct{}

func (s QiitaEntryAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Qiita Entry Request API ",
		"description": "Qiita Entry Request description",
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

func (s QiitaEntryAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Qiita Entry Response API",
		"description": "Qiita Entry Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"qiita_entries": map[string]interface{}{
				"type":  "array",
				"items": qiitaentry.Schema(),
			},
		},
		"required": []string{
			"qiita_entries",
		},
		"additionalProperties": false,
	}
	return schema
}
