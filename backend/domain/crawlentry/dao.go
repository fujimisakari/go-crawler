package crawlentry

import (
	"github.com/fujimisakari/go-crawler/backend/infrastracture/db"
)

type DAO interface {
	Select(limit int) ([]map[string]interface{}, error)
	SelectByDate(date string) (map[string]interface{}, error)
	Create(date string) (int, error)
}

func newDAO() DAO {
	return &db.CrawlEntryDAO{Cnn: db.SharedInstance().Connection}
}
