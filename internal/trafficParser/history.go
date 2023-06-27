package parser

import (
	"Wiggumize/utils"
)

type BrowseHistory struct {
	RequestsList []HistoryItem
	ListOfHosts  utils.Set
}

type HistoryItem struct {
	Time           string
	URL            string
	Host           string
	Path           string
	Method         string
	Request        string
	ReqHeaders     string
	ReqBody        string
	Status         string
	ReqContentType string
	Response       string
	ResHeaders     string
	ResBody        string
	ResContentType string
	Params         string
}

func (b *BrowseHistory) FilterByHost(hosts utils.Set) {
	filteredItems := []HistoryItem{}

	for _, item := range b.RequestsList {

		if hosts.Contains(item.Host) {
			filteredItems = append(filteredItems, item)
		}
	}
	b.RequestsList = filteredItems
	b.ListOfHosts = hosts
}
