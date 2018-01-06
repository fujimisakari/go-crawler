package service

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/builder"
	"github.com/fujimisakari/go-crawler/backend/domain/qiita_entry/entity"
)

func CrawlQiitaEntry(crawlEntryID int) (*entity.QiitaEntryList, error) {
	crawlerEngine := builder.QiitaEntryBuilder{}.BuildCrawlerEngine()

	fmt.Printf("Start: %s\n", crawlerEngine.Name)
	entries, err := crawlerEngine.Run()
	if err != nil {
		return nil, errors.Wrap(err, "qiita entry crawl error")
	}

	entities := []*entity.QiitaEntry{}
	for _, entry := range entries {
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
