package entity

import (
	"time"
)

type QiitaEntry struct {
	ID           int
	CrawlEntryID int
	User         string
	Title        string
	Link         string
	PostedAt     time.Time
}

func NewQiitaEntry(params map[string]interface{}) *QiitaEntry {
	return &QiitaEntry{
		ID:           params["id"].(int),
		CrawlEntryID: params["crawlEntryID"].(int),
		User:         params["user"].(string),
		Title:        params["title"].(string),
		Link:         params["link"].(string),
		PostedAt:     params["postedAt"].(time.Time),
	}
}

func (e *QiitaEntry) ToSchemaData() map[string]interface{} {
	layout := "2006-01-02 15:04:05"
	schemaData := map[string]interface{}{
		"id":        e.ID,
		"user":      e.User,
		"title":     e.Title,
		"link":      e.Link,
		"posted_at": e.PostedAt.Format(layout),
	}
	return schemaData
}
