package search

import (
	"Wiggumize/cli"
	parser "Wiggumize/internal/trafficParser"
	"fmt"
)

type SearchParams struct {
	ReqMethod      []searchMatch
	ReqHeader      []searchMatch
	ReqContentType []searchMatch
	ReqBody        []searchMatch
	ResHeader      []searchMatch
	ResContentType []searchMatch
	ResBody        []searchMatch
}

type searchMatch struct {
	value    string
	negative bool
}

type Search struct {
	Config   SearchConfig
	Regexp   SearchParams
	PHistory *parser.BrowseHistory
	channel  chan parser.HistoryItem
	Found    []parser.HistoryItem
	// HelpMessage string
}

func (s *Search) cleanUp() {
	s.Found = nil
	s.Regexp = SearchParams{}
}

func (s *Search) InputHandler() {

	// TODO: add history adn up arrow

	s.cleanUp()

	fmt.Print("Regexp Search. Type \"menu\" to get Search menu or \"exit\" to exit \n")

	input := cli.GetString("Type search query: ")

	switch input {
	case "menu", "Menu", "MENU":
		handleMenu(s)
	case "exit", "Exit", "EXIT":
		return
	case "":
	default:
		s.doSearch(input)
		s.output()
	}

	s.InputHandler()
}
