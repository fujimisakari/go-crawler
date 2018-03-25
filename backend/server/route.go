package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/fujimisakari/go-crawler/backend/handler"
	"github.com/fujimisakari/go-crawler/backend/jsonschema"
)

func setRoute(e *echo.Echo) {
	// Set Bundle MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.HTTPErrorHandler = handler.JSONHTTPError

	// Routes
	apiGroup := e.Group("/api")
	// Set Custom MiddleWare
	apiGroup.Use(jsonschema.Handler())
	{
		apiGroup.GET("/crawl_entry", handler.CrawlEntry)
		apiGroup.GET("/crawl/:date", handler.CrawlAggregate)
		apiGroup.GET("/hatena_hotentry/:date", handler.HatenaHotEntryList)
		apiGroup.GET("/hatena_hotentry_detail/:id", handler.HatenaHotEntryDetail)
		apiGroup.GET("/qiita_entry/:date", handler.QiitaEntryList)
		apiGroup.GET("/qiita_entry_detail/:id", handler.QiitaEntryDetail)
	}
}
