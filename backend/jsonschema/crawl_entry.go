package jsonschema

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
)

type CrawlEntryAPISchema struct{}

func (s CrawlEntryAPISchema) getRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":                "Crawl Entry Request API ",
		"description":          "Crawl Entry Request description",
		"method":               "GET",
		"type":                 "object",
		"properties":           map[string]interface{}{},
		"required":             []string{},
		"additionalProperties": false,
	}
	return schema
}

func (s CrawlEntryAPISchema) getResponseSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Crawl Entry Response API",
		"description": "Crawl Entry Response description",
		"type":        "object",
		"properties": map[string]interface{}{
			"crawl_entries": map[string]interface{}{
				"type":  "array",
				"items": crawlentry.Schema(),
			},
		},
		"required": []string{
			"crawl_entries",
		},
		"additionalProperties": false,
	}
	return schema
}
