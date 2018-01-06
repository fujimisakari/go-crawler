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

func (q *QiitaEntryList) PrintOut() {
	for _, item := range q.Items {
		fmt.Println("===============================")
		fmt.Println(item.PostedAt)
		fmt.Println(item.User)
		fmt.Println(item.Title)
		fmt.Println(item.Link)
	}
}
