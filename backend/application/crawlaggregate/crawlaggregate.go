package crawlaggregate

import (
	"github.com/fujimisakari/go-crawler/backend/domain/crawlaggregate/entity"
	"github.com/fujimisakari/go-crawler/backend/domain/crawlaggregate/service"
)

func Get(crawlDate string) (*entity.CrawlAggregate, error) {
	return service.GetCrawlAggregate(crawlDate)
}
