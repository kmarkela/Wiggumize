package search

import (
	"Wiggumize/cli"
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
	Config SearchConfig
	Regexp SearchParams
	// HelpMessage string
}

func hello() {
	msg := "Regexp Search. Type \"menu\" to get Search menu or \"exit\" to exit \n"
	// msg += "Type search quarry: "

	fmt.Print(msg)

}

func (s *Search) InputHandler() {

	hello()

	input := cli.GetString("Type search quarry: ")

	switch input {
	case "menu", "Menu", "MENU":
		handleMenu()
	case "exit", "Exit", "EXIT":
		return
	default:
		doSearch()
	}

	s.InputHandler()
}

func doSearch() {
	panic("unimplemented")
}
