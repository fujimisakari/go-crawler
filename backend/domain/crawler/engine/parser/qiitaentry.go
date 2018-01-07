package parser

import (
	"strings"
	"time"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
)

type QiitaEntryLogic struct{}

func (l QiitaEntryLogic) parse(res *response.Response) ([]map[string]interface{}, error) {
	items, err := res.ToJson()
	if err != nil {
		return nil, err
	}

	entries := []map[string]interface{}{}
	for _, item := range items {
		user := item["user"].(map[string]interface{})
		entry := map[string]interface{}{}
		entry["user"] = user["id"]
		entry["title"] = item["title"]
		entry["link"] = item["url"]
		entry["postedAt"] = toTime(item["created_at"].(string))
		entries = append(entries, entry)
	}
	return entries, nil
}

func toTime(date string) time.Time {
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "+09:00", "", -1)
	dateTime, _ := time.Parse("2006-01-02 15:04:05", date)
	return dateTime
}
