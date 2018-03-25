package result

type CrawlResult struct {
	Entries []map[string]interface{}
	Err     error
}

func NewChanel() chan *CrawlResult {
	return make(chan *CrawlResult)
}

func New(entries []map[string]interface{}, err error) *CrawlResult {
	return &CrawlResult{
		Entries: entries,
		Err:     err,
	}
}
