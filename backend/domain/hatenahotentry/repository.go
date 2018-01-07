package hatenahotentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

type Repository struct {
	dao DAO
}

func NewRepository() *Repository {
	return &Repository{dao: newDAO()}
}

func (r *Repository) FindByID(hatenaHotEntryID int) (*entity.HatenaHotEntry, error) {
	row, err := r.dao.SelectByID(hatenaHotEntryID)
	if err != nil {
		return nil, err
	}
	return toEntity(row), nil
}

func (r *Repository) FindByCrawlEntryID(crawlEntryID int) (*entity.HatenaHotEntryList, error) {
	rows, err := r.dao.SelectByCrawlEntryID(crawlEntryID)
	if err != nil {
		return nil, err
	}

	entitis := []*entity.HatenaHotEntry{}
	for _, row := range rows {
		entitis = append(entitis, toEntity(row))

	}
	return entity.NewHatenaHotEntryList(entitis), nil
}

func (r *Repository) BulkCreate(hatenaHotEntryList *entity.HatenaHotEntryList) error {
	err := r.dao.BulkInsert(hatenaHotEntryList)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteByCrawlEntryID(crawlEntryID int) error {
	err := r.dao.DeleteByCrawlEntryID(crawlEntryID)
	if err != nil {
		return err
	}
	return nil
}

func toEntity(row map[string]interface{}) *entity.HatenaHotEntry {
	params := map[string]interface{}{
		"id":           row["id"],
		"crawlEntryID": row["crawlEntryID"],
		"title":        row["title"],
		"link":         row["link"],
		"description":  row["description"],
	}
	return entity.NewHatenaHotEntry(params)
}
