package crawler

type Crawler interface {
	Crawl()
}

var mapper = map[string]interface{}{
	"hatena_hotentry": HatenaHotEntryCrawler{},
	"qiita_entry":     QiitaEntryCrawler{},
}

func GetCmdMapper() map[string]interface{} {
	return mapper
}

func GetCmdNameList() []string {
	cmdNameList := []string{}
	for k, _ := range mapper {
		cmdNameList = append(cmdNameList, k)
	}
	return cmdNameList
}
