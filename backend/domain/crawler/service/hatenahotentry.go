package service

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/result"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

func CrawlHatenaHotEntry(crawlEntryID int) (*entity.HatenaHotEntryList, error) {
	crawler := engine.NewCrawlerBuilder().HatenaHotEntry()
	done := result.NewChanel()
	defer close(done)

	fmt.Printf("Start: %s\n", crawler.Name)
	go crawler.Run(done)
	result := <-done
	if result.Err != nil {
		return nil, errors.Wrap(result.Err, "hatena hotentry crawl error")
	}

	entities := []*entity.HatenaHotEntry{}
	for _, entry := range result.Entries {
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
