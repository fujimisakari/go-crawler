package fetcher

import (
	"io/ioutil"
	"net/http"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
	"github.com/pkg/errors"
)

type FetchEngineLogic interface {
	getHttpMethod() string
	getCrawlUrl() string
	getHeaders() map[string]string
	isApiLimitExceeded() int
}

func New(logic FetchEngineLogic) *FetchEngine {
	return &FetchEngine{
		logic: logic,
	}
}

type FetchEngine struct {
	logic FetchEngineLogic
}

func (f FetchEngine) Fetch(res chan<- *response.Response) {
	method := f.logic.getHttpMethod()
	url := f.logic.getCrawlUrl()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		res <- response.New(nil, errors.Wrap(err, "fetch request create error"))
		return
	}
	headers := f.logic.getHeaders()
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		res <- response.New(nil, errors.Wrap(err, "fetch request error"))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res <- response.New(nil, errors.Wrap(err, "fetch response read error"))
		return
	}
	res <- response.New(body, nil)
}
