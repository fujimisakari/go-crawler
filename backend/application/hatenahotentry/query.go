package hatenahotentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry"
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

func GetByID(id int) (*entity.HatenaHotEntry, error) {
	hatenaHotEntryRepo := hatenahotentry.NewRepository()
	hatenaHotEntry, err := hatenaHotEntryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return hatenaHotEntry, nil
}

func GetListByDate(crawlDate string) (*entity.HatenaHotEntryList, error) {
	crawlEntryRepo := crawlentry.NewRepository()
	crawlEntry, err := crawlEntryRepo.FindOrCreateByDate(crawlDate)
	if err != nil {
		return nil, err
	}

	hatenaHotEntryRepo := hatenahotentry.NewRepository()
	hatenaHotEntryList, err := hatenaHotEntryRepo.FindByCrawlEntryID(crawlEntry.ID)
	if err != nil {
		return nil, err
	}
	return hatenaHotEntryList, nil
}
