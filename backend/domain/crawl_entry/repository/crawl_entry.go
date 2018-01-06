package reposiotry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawl_entry"
	"github.com/fujimisakari/go-crawler/backend/domain/crawl_entry/entity"
	"github.com/fujimisakari/go-crawler/backend/registry/dao"
)

type CrawlEntryRepository struct {
	dao crawl_entry.CrawlEntryDAO
}

func New() *CrawlEntryRepository {
	return &CrawlEntryRepository{dao: dao.New().CrawlEntry()}
}

func (r *CrawlEntryRepository) Find(limit int) (*entity.CrawlEntryList, error) {
	rows, err := r.dao.Select(limit)
	if err != nil {
		return nil, err
	}

	entitis := []*entity.CrawlEntry{}
	for _, row := range rows {
		entitis = append(entitis, toEntity(row["id"].(int), row["date"].(string)))
	}
	return entity.NewCrawlEntryList(entitis), nil
}

func (r *CrawlEntryRepository) FindOrCreateByDate(date string) (*entity.CrawlEntry, error) {
	// Find
	row, err := r.dao.SelectByDate(date)
	if err != nil {
		return nil, err
	}
	if row != nil {
		return entity.NewCrawlEntry(row), nil
	}

	// Create
	id, err := r.dao.Create(date)
	if err != nil {
		return nil, err
	}
	return toEntity(id, date), nil
}

func toEntity(id int, date string) *entity.CrawlEntry {
	params := map[string]interface{}{
		"id":   id,
		"date": date,
	}
	return entity.NewCrawlEntry(params)
}
