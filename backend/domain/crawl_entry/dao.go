package crawl_entry

type CrawlEntryDAO interface {
	Select(limit int) ([]map[string]interface{}, error)
	SelectByDate(date string) (map[string]interface{}, error)
	Create(date string) (int, error)
}
