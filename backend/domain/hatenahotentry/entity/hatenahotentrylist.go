package entity

import (
	"fmt"
)

type HatenaHotEntryList struct {
	Items []*HatenaHotEntry
}

func NewHatenaHotEntryList(items []*HatenaHotEntry) *HatenaHotEntryList {
	return &HatenaHotEntryList{
		Items: items,
	}
}

func (e *HatenaHotEntryList) ToSchemaData() []map[string]interface{} {
	schemaDatas := []map[string]interface{}{}
	for _, item := range e.Items {
		schemaDatas = append(schemaDatas, item.ToSchemaData())
	}
	return schemaDatas
}

func (e *HatenaHotEntryList) PrintOut() {
	for _, item := range e.Items {
		fmt.Println("===============================")
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
	}
}
