package search

import (
	"Wiggumize/utils"
	"fmt"
)

type SearchParams struct {
	ReqMethod      []searchMatch
	ReqHeader      []searchMatch
	ReqContentType []searchMatch
	ReqBody        []searchMatch
	ResMethod      []searchMatch
	ResHeader      []searchMatch
	ResContentType []searchMatch
	ResBody        []searchMatch
}

type searchMatch map[matchType]string

type matchType string

const (
	contain    matchType = "contain"
	eq         matchType = "eq"
	notEq      matchType = "notEq"
	notContain matchType = "notContain"
)

type SearchConfig struct {
	Output outputType
}

type outputType int

const (
	endpoint outputType = iota
	headers
	reqOnly
	full
)

type Search struct {
	Config      SearchConfig
	Regexp      SearchParams
	HelpMessage string
}

func (s *Search) Help() {
	fmt.Println(s.HelpMessage)
}

func returnHelpMsg() string {
	help := "\n\nRegexp Search: \n\n"
	help += "Avaliable search fields: \n"

	searchFields := utils.GetStructFieldNames(SearchParams{})

	for _, name := range searchFields {
		help += "- " + name + "\n"
	}

	help += "\nAvaliable search operators: \n"

	help += "- contain\n"
	help += "- eq\n"
	help += "- notEq\n"
	help += "- notContain\n\n"

	help += "Search Example: \n"
	help += "ReqMethod.eq POST & ReqBody.contain admin & ResContentType.notContain HTML & ResBody.contain success\n"

	return help
}

func BuildSearch() Search {

	return Search{
		HelpMessage: returnHelpMsg(),
	}
}
