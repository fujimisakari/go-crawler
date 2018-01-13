package crawlentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry/entity"
)

func GetList() (*entity.CrawlEntryList, error) {
	crawlEntryRepo := crawlentry.NewRepository()
	return crawlEntryRepo.Find(10)
}
