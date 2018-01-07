package fetcher

type HatenaHotEntryLogic struct{}

func (l HatenaHotEntryLogic) getHttpMethod() string {
	return "GET"
}

func (l HatenaHotEntryLogic) getCrawlUrl() string {
	return "http://b.hatena.ne.jp/hotentry?mode=rss"
}

func (l HatenaHotEntryLogic) getHeaders() map[string]string {
	return map[string]string{}
}

func (l HatenaHotEntryLogic) isApiLimitExceeded() int {
	return 5
}
