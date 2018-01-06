package parser

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"

	"github.com/fujimisakari/go-crawler/backend/domain/crawler/engine/response"
)

type HatenaHotEntryLogic struct{}

func (l HatenaHotEntryLogic) parse(res *response.Response) ([]map[string]interface{}, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(res.ToString())
	if err != nil {
		return nil, errors.Wrap(err, "feed parse error")
	}

	entries := []map[string]interface{}{}
	for _, item := range feed.Items {
		entry := map[string]interface{}{}
		entry["title"] = item.Title
		entry["link"] = item.Link
		entry["description"] = item.Description
		entries = append(entries, entry)
	}
	return entries, nil
}
