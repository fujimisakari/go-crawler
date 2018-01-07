package entity

type CrawlEntryList struct {
	Items []*CrawlEntry
}

func NewCrawlEntryList(items []*CrawlEntry) *CrawlEntryList {
	return &CrawlEntryList{
		Items: items,
	}
}
