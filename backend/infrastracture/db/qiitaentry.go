package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

type QiitaEntryDAO struct {
	Cnn *sql.DB
}

func (r *QiitaEntryDAO) SelectByID(qiitaEntryID int) (map[string]interface{}, error) {
	var id int
	var crawlEntryID int
	var user string
	var title string
	var link string
	var postedAt time.Time
	query := "SELECT id, crawl_entry_id, user, title, link, posted_at FROM qiita_entry WHERE id = ?;"
	err := r.Cnn.QueryRow(query, qiitaEntryID).Scan(&id, &crawlEntryID, &user, &title, &link, &postedAt)
	if err != nil {
		return nil, errors.Wrap(err, "sql select error")
	}

	row := map[string]interface{}{
		"id":           id,
		"crawlEntryID": crawlEntryID,
		"user":         user,
		"title":        title,
		"link":         link,
		"postedAt":     postedAt,
	}
	return row, nil
}

func (r *QiitaEntryDAO) SelectByCrawlEntryID(crawlEntryID int) ([]map[string]interface{}, error) {
	query := "SELECT id, crawl_entry_id, user, title, link, posted_at FROM qiita_entry WHERE crawl_entry_id = ?;"
	result, err := r.Cnn.Query(query, crawlEntryID)
	if err != nil {
		return nil, errors.Wrap(err, "sql select error")
	}
	defer result.Close()

	rows := []map[string]interface{}{}
	for result.Next() {
		var id int
		var crawlEntryID int
		var user string
		var title string
		var link string
		var postedAt time.Time
		_ = result.Scan(&id, &crawlEntryID, &user, &title, &link, &postedAt)
		row := map[string]interface{}{
			"id":           id,
			"crawlEntryID": crawlEntryID,
			"title":        title,
			"link":         link,
			"postedAt":     postedAt,
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *QiitaEntryDAO) BulkInsert(qiitaEntryList *entity.QiitaEntryList) error {
	lastIdx := len(qiitaEntryList.Items) - 1
	query := "INSERT INTO qiita_entry (crawl_entry_id, user, title, link, posted_at) VALUES "
	for idx, item := range qiitaEntryList.Items {
		queryFmt := "('%d', '%s', '%s', '%s', '%s')"
		query += fmt.Sprintf(queryFmt, item.CrawlEntryID, escape(item.User), escape(item.Title), item.Link, toStrTime(item.PostedAt))
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

func (r *QiitaEntryDAO) DeleteByCrawlEntryID(crawlEntryID int) error {
	_, err := r.Cnn.Exec("DELETE FROM qiita_entry WHERE crawl_entry_id = ?", crawlEntryID)
	if err != nil {
		return errors.Wrap(err, "sql delete error")
	}
	return nil
}
