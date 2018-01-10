package entity

import (
	hatenahotentry "github.com/fujimisakari/go-crawler/backend/domain/hatenahotentry/entity"
	qiitaentry "github.com/fujimisakari/go-crawler/backend/domain/qiitaentry/entity"
)

type CrawlAggregate struct {
	Date               string
	HatenaHotEntryList *hatenahotentry.HatenaHotEntryList
	QiitaEntryList     *qiitaentry.QiitaEntryList
}

func NewCrawlAggregate(params map[string]interface{}) *CrawlAggregate {
	return &CrawlAggregate{
		Date:               params["date"].(string),
		HatenaHotEntryList: params["hatenaHotEntryList"].(*hatenahotentry.HatenaHotEntryList),
		QiitaEntryList:     params["qiitaEntryList"].(*qiitaentry.QiitaEntryList),
	}
}

func (e *CrawlAggregate) ToSchemaData() map[string]interface{} {
	schemaData := map[string]interface{}{
		"date":              e.Date,
		"hatena_hotentries": e.HatenaHotEntryList.ToSchemaData(),
		"qiita_entries":     e.QiitaEntryList.ToSchemaData(),
	}
	return schemaData
}
