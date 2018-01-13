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

func (e *CrawlEntry) ToSchemaData() map[string]interface{} {
	schemaData := map[string]interface{}{
		"id":    e.ID,
		"title": e.Date,
	}
	return schemaData
}
