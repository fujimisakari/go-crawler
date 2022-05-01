package crawler

import (
	"log"
	"time"

	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/service"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
)

type QiitaEntryCrawler struct{}

func (c QiitaEntryCrawler) Crawl() {
	crawlEntryRepo := crawlentry.NewRepository()
	qiitaEntryRepo := qiitaentry.NewRepository()

	crawlDate := time.Now().Format("2006-01-02")
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Crawling
	qiitaEntryList, err := service.CrawlQiitaEntry(crawlEntry.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Save
	err = qiitaEntryRepo.DeleteByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = qiitaEntryRepo.BulkCreate(qiitaEntryList)
	if err != nil {
		log.Fatal(err)
		return
	}

	qiitaEntryList.PrintOut()
}
