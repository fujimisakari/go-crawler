package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
)

type CrawlAggregateAPISchema struct{}

func (s CrawlAggregateAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Crawl Aggregate Request API ",
		"description": "Crawl Aggregate Request description",
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

func (s CrawlAggregateAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Crawl Aggregate Response API",
		"description": "Crawl Aggregate Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"date": map[string]interface{}{
				"type":        "string",
				"example":     "2018-1-1",
				"description": "crawl date",
			},
			"hatena_hotentries": map[string]interface{}{
				"type":  "array",
				"items": hatenahotentry.Schema(),
			},
			"qiita_entries": map[string]interface{}{
				"type":  "array",
				"items": qiitaentry.Schema(),
			},
		},
		"required": []string{
			"date",
			"hatena_hotentries",
			"qiita_entries",
		},
		"additionalProperties": false,
	}
	return schema
}
