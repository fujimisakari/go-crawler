package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/fujimisakari/go-crawler/backend/application/hatenahotentry"
)

func HatenaHotEntryList(c echo.Context) error {
	date := c.Param("date")
	return c.String(http.StatusOK, date)
}

func HatenaHotEntryDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	hatenaHotEntry, err := hatenahotentry.GetByID(id)
	if err != nil {
		return err
	}

	context := map[string]interface{}{
		"hatena_hotentry": hatenaHotEntry.ToSchemaData(),
	}
	c.Set("context", context)
	return nil
}

func QiitaEntry(c echo.Context) error {
	date := c.Param("date")
	return c.String(http.StatusOK, date)
}
