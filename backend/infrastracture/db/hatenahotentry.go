package db

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
)

type HatenaHotEntryDAO struct {
	Cnn *sql.DB
}

func (r *HatenaHotEntryDAO) SelectByID(hatenaHotEntryID int) (map[string]interface{}, error) {
	var id int
	var crawlEntryID int
	var title string
	var link string
	var description string
	query := "SELECT id, crawl_entry_id, title, link, description FROM hatena_hotentry WHERE id = ?;"
	err := r.Cnn.QueryRow(query, hatenaHotEntryID).Scan(&id, &crawlEntryID, &title, &link, &description)
	if err != nil {
		return nil, errors.Wrap(err, "sql select error")
	}

	row := map[string]interface{}{
		"id":           id,
		"crawlEntryID": crawlEntryID,
		"title":        title,
		"link":         link,
		"description":  description,
	}
	return row, nil
}

func (r *HatenaHotEntryDAO) SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error) {
	query := "SELECT id, crawl_entry_id, title, link, description FROM hatena_hotentry WHERE crawl_entry_id = ?;"
	result, err := r.Cnn.Query(query, crawlEntryID)
	if err != nil {
		return nil, errors.Wrap(err, "sql select error")
	}
	defer result.Close()

	rows := []map[string]interface{}{}
	for result.Next() {
		var id int
		var crawlEntryID int
		var title string
		var link string
		var description string
		_ = result.Scan(&id, &crawlEntryID, &title, &link, &description)
		row := map[string]interface{}{
			"id":           id,
			"crawlEntryID": crawlEntryID,
			"title":        title,
			"link":         link,
			"description":  description,
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *HatenaHotEntryDAO) BulkInsert(hatenaHotEntryList *entity.HatenaHotEntryList) error {
	lastIdx := len(hatenaHotEntryList.Items) - 1
	query := "INSERT INTO hatena_hotentry (crawl_entry_id, title, link, description) VALUES "
	for idx, item := range hatenaHotEntryList.Items {
		query += fmt.Sprintf("('%d', '%s', '%s', '%s')", item.CrawlEntryID, escape(item.Title), item.Link, escape(item.Description))
		if idx == lastIdx {
			query += ";"
		} else {
			query += ","
		}
	}

	_, err := r.Cnn.Exec(query)
	if err != nil {
		return errors.Wrap(err, "sql bulk create error")
	}
	return nil
}

func (r *HatenaHotEntryDAO) DeleteByCrawlEntryID(crawlEntryID int) error {
	_, err := r.Cnn.Exec("DELETE FROM hatena_hotentry WHERE crawl_entry_id = ?", crawlEntryID)
	if err != nil {
		return errors.Wrap(err, "sql delete error")
	}
	return nil
}
