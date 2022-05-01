package jsonschema

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

type APISchema interface {
	getRequestSchema() map[string]interface{}
	getResponseSchema() map[string]interface{}
}

func Handler() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			// Define schema
			schemaMapper := getSchemaMapper()
			urlMethod := strings.Replace(c.Path(), "/api/", "", 1)
			apiSchema, ok := schemaMapper[urlMethod].(APISchema)
			if !ok {
				logrus.Error("Dose not exsit SchemaMapper: ", c.Path())
				jsonError := jsonError(http.StatusNotFound)
				return c.JSON(http.StatusNotFound, jsonError)
			}

			// Request validate check
			reqSchema := apiSchema.getRequestSchema()
			reqContext := createRequestContext(c)
			if result, msg := validate(reqSchema, reqContext); result != true {
				logrus.Error("Request validate error")
				fmt.Printf("%+v\n", msg)
				jsonError := jsonError(http.StatusForbidden)
				return c.JSON(http.StatusForbidden, jsonError)
			}

			// Execute echo handler
			if err := next(c); err != nil {
				jsonError := jsonError(http.StatusInternalServerError)
				jsonError["message"] = err.Error()
				fmt.Printf("%+v\n", err)
				return c.JSON(http.StatusInternalServerError, jsonError)
			}

			// Response validate check
			resSchema := apiSchema.getResponseSchema()
			resContext := c.Get("context")
			if result, msg := validate(resSchema, resContext); result != true {
				logrus.Error("Response validate error")
				fmt.Printf("%+v\n", msg)
				jsonError := jsonError(http.StatusInternalServerError)
				return c.JSON(http.StatusInternalServerError, jsonError)
			}

			return c.JSON(http.StatusOK, resContext)
		})
	}
}

func validate(schema interface{}, context interface{}) (bool, string) {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	contextLoader := gojsonschema.NewGoLoader(context)

	result, err := gojsonschema.Validate(schemaLoader, contextLoader)
	if err != nil {
		logrus.Error("JsonSchema Validate Error:", err)
	}

	var msg string
	if !result.Valid() {
		msg += "The Json is not valid. see errors"
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			msg += fmt.Sprintf("\n- %s", err)
		}
	}
	return result.Valid(), msg
}

func createRequestContext(c echo.Context) map[string]interface{} {
	mapList := make(map[string]interface{})

	switch c.Request().Method {
	case echo.GET:
		for _, n := range c.ParamNames() {
			val, err := strconv.ParseInt(c.Param(n), 0, 64)
			if err != nil {
				mapList[n] = c.Param(n)
			} else {
				mapList[n] = val
			}
		}
	case echo.POST:
	case echo.PUT:
	case echo.DELETE:
	case echo.PATCH:
	case echo.OPTIONS:
	case echo.HEAD:
	case echo.CONNECT:
	case echo.TRACE:
	}
	return mapList
}

func jsonError(statusCode int) map[string]string {
	error := map[string]string{
		"statusCode": strconv.Itoa(statusCode),
		"message":    http.StatusText(statusCode),
	}
	return error
}
