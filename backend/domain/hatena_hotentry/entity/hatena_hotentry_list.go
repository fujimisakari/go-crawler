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

func (h *HatenaHotEntryList) PrintOut() {
	for _, item := range h.Items {
		fmt.Println("===============================")
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
	}
}
