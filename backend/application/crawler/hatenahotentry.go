package crawler

import (
	"log"
	"time"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/service"

	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
)

type HatenaHotEntryCrawler struct{}

func (c HatenaHotEntryCrawler) Crawl() {
	crawlEntryRepo := crawlentry.NewRepository()
	hatenaHotEntryRepo := hatenahotentry.NewRepository()

	crawlDate := time.Now().Format("2006-01-02")
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Crawling
	hatenaHotEntryList, err := service.CrawlHatenaHotEntry(crawlEntry.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Save
	err = hatenaHotEntryRepo.DeleteByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = hatenaHotEntryRepo.BulkCreate(hatenaHotEntryList)
	if err != nil {
		log.Fatal(err)
		return
	}
	hatenaHotEntryList.PrintOut()
}
