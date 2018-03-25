package handler

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/fujimisakari/go-crawler/backend/application/qiitaentry"
)

func QiitaEntryList(c echo.Context) error {
	date := c.Param("date")
	qiitaEntryList, err := qiitaentry.GetListByDate(date)
	if err != nil {
		return err
	}

	context := map[string]interface{}{
		"qiita_entries": qiitaEntryList.ToSchemaData(),
	}
	c.Set("context", context)
	return nil
}

func QiitaEntryDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	qiitaEntry, err := qiitaentry.GetByID(id)
	if err != nil {
		return err
	}

	context := map[string]interface{}{
		"qiita_entry": qiitaEntry.ToSchemaData(),
	}
	c.Set("context", context)
	return nil
}
