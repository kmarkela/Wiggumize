package search

import (
	"fmt"
	"strings"
	"unicode"
)

func parseNot(input string) (string, searchMatch) {

	var sm searchMatch = searchMatch{}
	match := matchType(1)
	input = strings.TrimLeftFunc(input, unicode.IsSpace)

	// TODO: refactor this
	if strings.HasPrefix(input, "!") {
		match = matchType(0)
		// Remove !
		input = strings.Split(input, "!")[1]
		// remove spaces
		input = strings.TrimLeftFunc(input, unicode.IsSpace)
	}
	var p []string = strings.Split(input, " ")

	// ToDo: error handling. check if len(p) > 1
	input = strings.Join(p[1:], " ")

	sm[match] = input

	return p[0], sm
}

func parseInput(input string) (error, SearchParams) {

	var parseAnd []string
	var sp SearchParams = SearchParams{}

	parseAnd = strings.Split(input, "&")

	for _, v := range parseAnd {

		k, match := parseNot(v)
		fmt.Println(k)
		switch k {
		case "ReqMethod":
			sp.ReqMethod = append(sp.ReqMethod, match)
		}

	}

	return nil, sp
}

func (s *Search) doSearch(input string) {

	_, s.Regexp = parseInput(input)
	fmt.Println(s.Regexp)

}
