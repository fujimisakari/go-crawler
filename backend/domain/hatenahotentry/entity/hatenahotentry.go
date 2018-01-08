package entity

type HatenaHotEntry struct {
	ID           int
	CrawlEntryID int
	Title        string
	Link         string
	Description  string
}

func NewHatenaHotEntry(params map[string]interface{}) *HatenaHotEntry {
	return &HatenaHotEntry{
		ID:           params["id"].(int),
		CrawlEntryID: params["crawlEntryID"].(int),
		Title:        params["title"].(string),
		Link:         params["link"].(string),
		Description:  params["description"].(string),
	}
}

func (e *HatenaHotEntry) ToSchemaData() map[string]interface{} {
	schemaData := map[string]interface{}{
		"id":          e.ID,
		"title":       e.Title,
		"link":        e.Link,
		"description": e.Description,
	}
	return schemaData
}
