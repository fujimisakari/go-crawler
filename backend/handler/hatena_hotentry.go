package handler

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/fujimisakari/go-crawler/backend/application/hatenahotentry"
)

func HatenaHotEntryList(c echo.Context) error {
	date := c.Param("date")
	hatenaHotEntryList, err := hatenahotentry.GetListByDate(date)
	if err != nil {
		return err
	}

	context := map[string]interface{}{
		"hatena_hotentries": hatenaHotEntryList.ToSchemaData(),
	}
	c.Set("context", context)
	return nil
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
