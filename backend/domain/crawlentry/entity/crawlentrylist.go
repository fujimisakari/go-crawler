package entity

type CrawlEntryList struct {
	Items []*CrawlEntry
}

func NewCrawlEntryList(items []*CrawlEntry) *CrawlEntryList {
	return &CrawlEntryList{
		Items: items,
	}
}

func (e *CrawlEntryList) ToSchemaData() []map[string]interface{} {
	schemaDatas := []map[string]interface{}{}
	for _, item := range e.Items {
		crawlEntry := item.ToSchemaData()
		crawlEntry["crawl_items"] = []string{"hatena_hotentry", "qiita_entry"}
		schemaDatas = append(schemaDatas, crawlEntry)
	}
	return schemaDatas
}
