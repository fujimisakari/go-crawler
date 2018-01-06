package crawler

import (
	"log"
	"time"

	crawler_sv "github.com/fujimisakari/go-crawler/backend/domain/crawler/service"

	crawl_entry_repo "github.com/fujimisakari/go-crawler/backend/domain/crawl_entry/repository"
	qiita_entry_repo "github.com/fujimisakari/go-crawler/backend/domain/qiita_entry/repository"
)

type QiitaEntryCrawler struct{}

func (c QiitaEntryCrawler) Crawl() {
	crawlEntryRepo := crawl_entry_repo.New()
	qiitaEntryRepo := qiita_entry_repo.New()

	crawlDate := time.Now().Format("2006-01-02")
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Crawling
	qiitaEntryList, err := crawler_sv.CrawlQiitaEntry(crawlEntry.ID)
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
