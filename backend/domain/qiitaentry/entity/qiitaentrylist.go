package entity

import (
	"fmt"
)

type QiitaEntryList struct {
	Items []*QiitaEntry
}

func NewQiitaEntryList(items []*QiitaEntry) *QiitaEntryList {
	return &QiitaEntryList{
		Items: items,
	}
}

func (e *QiitaEntryList) ToSchemaData() []map[string]interface{} {
	schemaDatas := []map[string]interface{}{}
	for _, item := range e.Items {
		schemaDatas = append(schemaDatas, item.ToSchemaData())
	}
	return schemaDatas
}

func (e *QiitaEntryList) PrintOut() {
	for _, item := range e.Items {
		fmt.Println("===============================")
		fmt.Println(item.PostedAt)
		fmt.Println(item.User)
		fmt.Println(item.Title)
		fmt.Println(item.Link)
	}
}
