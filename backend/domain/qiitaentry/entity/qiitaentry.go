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
