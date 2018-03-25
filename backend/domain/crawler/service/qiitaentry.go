package service

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/result"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

func CrawlQiitaEntry(crawlEntryID int) (*entity.QiitaEntryList, error) {
	crawler := engine.NewCrawlerBuilder().QiitaEntry()
	done := result.NewChanel()
	defer close(done)

	fmt.Printf("Start: %s\n", crawler.Name)
	go crawler.Run(done)
	result := <-done
	if result.Err != nil {
		return nil, errors.Wrap(result.Err, "qiita entry crawl error")
	}

	entities := []*entity.QiitaEntry{}
	for _, entry := range result.Entries {
		params := map[string]interface{}{
			"id":           0,
			"crawlEntryID": crawlEntryID,
			"user":         entry["user"],
			"title":        entry["title"],
			"link":         entry["link"],
			"postedAt":     entry["postedAt"],
		}
		entities = append(entities, entity.NewQiitaEntry(params))
	}
	return entity.NewQiitaEntryList(entities), nil
}
