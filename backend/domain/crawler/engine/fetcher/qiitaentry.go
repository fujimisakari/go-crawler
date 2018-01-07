package fetcher

type QiitaEntryLogic struct{}

func (l QiitaEntryLogic) getHttpMethod() string {
	return "GET"
}

func (l QiitaEntryLogic) getCrawlUrl() string {
	return "https://qiita.com/api/v2/items"
}

func (l QiitaEntryLogic) getHeaders() map[string]string {
	return map[string]string{}
}

func (l QiitaEntryLogic) isApiLimitExceeded() int {
	return 5
}
