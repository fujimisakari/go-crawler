package api

import (
	"github.com/labstack/echo"

	"github.com/fujimisakari/go-crawler/backend/application/crawlaggregate"
)

func CrawlAggregate(c echo.Context) error {
	date := c.Param("date")
	crawlAggregate, err := crawlaggregate.Get(date)
	if err != nil {
		return err
	}

	c.Set("context", crawlAggregate.ToSchemaData())
	return nil
}
