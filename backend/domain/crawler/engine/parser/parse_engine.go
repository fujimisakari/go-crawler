package parser

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
)

type ParseEngineLogic interface {
	parse(res *response.Response) ([]map[string]interface{}, error)
}

func New(logic ParseEngineLogic) *ParseEngine {
	return &ParseEngine{
		logic: logic,
	}
}

type ParseEngine struct {
	logic ParseEngineLogic
}

func (f ParseEngine) Parse(res *response.Response) ([]map[string]interface{}, error) {
	return f.logic.parse(res)
}
