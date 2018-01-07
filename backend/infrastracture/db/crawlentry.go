package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

type CrawlEntryDAO struct {
	Cnn *sql.DB
}

func (r *CrawlEntryDAO) Select(limit int) ([]map[string]interface{}, error) {
	result, err := r.Cnn.Query("SELECT id, date FROM crawl_entry ORDER BY id LIMIT ?;", limit)
	if err != nil {
		return nil, errors.Wrap(err, "sql select error")
	}
	defer result.Close()

	rows := []map[string]interface{}{}
	for result.Next() {
		var id int
		var date string
		err = result.Scan(&id, &date)
		if err != nil {
			return nil, errors.Wrap(err, "sql scan error")
		}
		row := map[string]interface{}{
			"id":   id,
			"date": date,
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (r *CrawlEntryDAO) SelectByDate(crawlDate string) (map[string]interface{}, error) {
	var id int
	var date string
	err := r.Cnn.QueryRow("SELECT id, date FROM crawl_entry WHERE date = ?;", crawlDate).Scan(&id, &date)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrap(err, "sql select error")
	}

	row := map[string]interface{}{
		"id":   id,
		"date": date,
	}
	return row, nil
}

func (r *CrawlEntryDAO) Create(date string) (int, error) {
	result, err := r.Cnn.Exec("INSERT INTO crawl_entry (date) VALUES (?);", date)
	if err != nil {
		return 0, errors.Wrap(err, "sql create error")
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}
