package hatena_hotentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatena_hotentry/entity"
)

type HatenaHotEntryDAO interface {
	SelectByID(id int) (map[string]interface{}, error)
	SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error)
	BulkInsert(hatenaHotEntryList *entity.HatenaHotEntryList) error
	DeleteByCrawlEntryID(crawlEntryID int) error
}
