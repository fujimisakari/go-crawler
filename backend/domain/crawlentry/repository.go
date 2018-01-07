package crawlentry

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlentry/entity"
)

type Repository struct {
	dao DAO
}

func NewRepository() *Repository {
	return &Repository{dao: newDAO()}
}

func (r *Repository) Find(limit int) (*entity.CrawlEntryList, error) {
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

func (r *Repository) FindOrCreateByDate(date string) (*entity.CrawlEntry, error) {
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
