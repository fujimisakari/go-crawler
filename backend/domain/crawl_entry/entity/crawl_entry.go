package entity

type CrawlEntry struct {
	ID   int
	Date string
}

func NewCrawlEntry(params map[string]interface{}) *CrawlEntry {
	return &CrawlEntry{
		ID:   params["id"].(int),
		Date: params["date"].(string),
	}
}
