package service

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

func CrawlHatenaHotEntry(crawlEntryID int) (*entity.HatenaHotEntryList, error) {
	crawler := engine.NewCrawlerBuilder().HatenaHotEntry()

	fmt.Printf("Start: %s\n", crawler.Name)
	entries, err := crawler.Run()
	if err != nil {
		return nil, errors.Wrap(err, "hatena hotentry crawl error")
	}

	entities := []*entity.HatenaHotEntry{}
	for _, entry := range entries {
		params := map[string]interface{}{
			"id":           0,
			"crawlEntryID": crawlEntryID,
			"title":        entry["title"],
			"link":         entry["link"],
			"description":  entry["description"],
		}
		entity := entity.NewHatenaHotEntry(params)
		entities = append(entities, entity)
	}
	return entity.NewHatenaHotEntryList(entities), nil
}
