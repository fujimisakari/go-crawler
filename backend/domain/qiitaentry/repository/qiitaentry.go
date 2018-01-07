package reposiotry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry"
	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
	"github.com/fujimisakari/go-crawler/backend/registry/dao"
)

type QiitaEntryRepository struct {
	dao qiitaentry.QiitaEntryDAO
}

func New() *QiitaEntryRepository {
	return &QiitaEntryRepository{dao: dao.New().QiitaEntry()}
}

func (r *QiitaEntryRepository) FindByID(hatenaHotEntryID int) (*entity.QiitaEntry, error) {
	row, err := r.dao.SelectByID(hatenaHotEntryID)
	if err != nil {
		return nil, err
	}
	return toEntity(row), nil
}

func (r *QiitaEntryRepository) FindByCrawlEntryID(crawlEntryID int) (*entity.QiitaEntryList, error) {
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

func (r *QiitaEntryRepository) BulkCreate(qiitaEntryList *entity.QiitaEntryList) error {
	err := r.dao.BulkInsert(qiitaEntryList)
	if err != nil {
		return err
	}
	return nil
}

func (r *QiitaEntryRepository) DeleteByCrawlEntryID(crawlEntryID int) error {
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
