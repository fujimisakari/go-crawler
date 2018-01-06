package crawler

import (
	"log"
	"time"

	crawler_sv "github.com/fujimisakari/go-crawler/backend/domain/crawler/service"

	crawl_entry_repo "github.com/fujimisakari/go-crawler/backend/domain/crawl_entry/repository"
	hatena_hotentry_repo "github.com/fujimisakari/go-crawler/backend/domain/hatena_hotentry/repository"
)

type HatenaHotEntryCrawler struct{}

func (c HatenaHotEntryCrawler) Crawl() {
	crawlEntryRepo := crawl_entry_repo.New()
	hatenaHotEntryRepo := hatena_hotentry_repo.New()

	crawlDate := time.Now().Format("2006-01-02")
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Crawling
	hatenaHotEntryList, err := crawler_sv.CrawlHatenaHotEntry(crawlEntry.ID)
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
