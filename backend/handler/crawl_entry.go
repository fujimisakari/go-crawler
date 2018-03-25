package handler

import (
	"github.com/labstack/echo"

	"github.com/fujimisakari/go-crawler/backend/application/crawlentry"
)

func CrawlEntry(c echo.Context) error {
	crawlEntryList, err := crawlentry.GetList()
	if err != nil {
		return err
	}

	context := map[string]interface{}{
		"crawl_entries": crawlEntryList.ToSchemaData(),
	}
	c.Set("context", context)
	return nil
}
