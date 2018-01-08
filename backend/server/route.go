package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/fujimisakari/go-crawler/backend/server/api"
	"github.com/fujimisakari/go-crawler/backend/server/jsonschema"
)

func setRoute(e *echo.Echo) {
	// Set Bundle MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.HTTPErrorHandler = JSONHTTPErrorHandler

	// Routes
	apiGroup := e.Group("/api")
	// Set Custom MiddleWare
	apiGroup.Use(jsonschema.JSONSchemaHandler())
	{
		apiGroup.GET("/hatena_hotentry/:date", api.HatenaHotEntryList)
		apiGroup.GET("/hatena_hotentry_detail/:id", api.HatenaHotEntryDetail)
		apiGroup.GET("/qitta_entry/:date", api.QiitaEntry)
	}
}
