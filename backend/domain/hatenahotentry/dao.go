package hatenahotentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

type HatenaHotEntryDAO interface {
	SelectByID(id int) (map[string]interface{}, error)
	SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error)
	BulkInsert(hatenaHotEntryList *entity.HatenaHotEntryList) error
	DeleteByCrawlEntryID(crawlEntryID int) error
}
