package qiitaentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

type Repository struct {
	dao DAO
}

func NewRepository() *Repository {
	return &Repository{dao: newDAO()}
}

func (r *Repository) FindByID(hatenaHotEntryID int) (*entity.QiitaEntry, error) {
	row, err := r.dao.SelectByID(hatenaHotEntryID)
	if err != nil {
		return nil, err
	}
	return toEntity(row), nil
}

func (r *Repository) FindByCrawlEntryID(crawlEntryID int) (*entity.QiitaEntryList, error) {
	rows, err := r.dao.SelectByCrawlEntryID(crawlEntryID)
	if err != nil {
		return nil, err
	}

	entitis := []*entity.QiitaEntry{}
	for _, row := range rows {
		entitis = append(entitis, toEntity(row))

	}
	return entity.NewQiitaEntryList(entitis), nil
}

func (r *Repository) BulkCreate(qiitaEntryList *entity.QiitaEntryList) error {
	err := r.dao.BulkInsert(qiitaEntryList)
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

func toEntity(row map[string]interface{}) *entity.QiitaEntry {
	params := map[string]interface{}{
		"id":           row["id"],
		"crawlEntryID": row["crawlEntryID"],
		"user":         row["user"],
		"title":        row["title"],
		"link":         row["link"],
		"postedAt":     row["postedAt"],
	}
	return entity.NewQiitaEntry(params)
}
