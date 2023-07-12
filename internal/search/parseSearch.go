package search

import (
	"fmt"
	"strings"
	"unicode"
)

func parseNot(input string) (string, searchMatch) {

	n := false
	// input = strings.TrimLeftFunc(input, unicode.IsSpace)
	input = strings.TrimFunc(input, unicode.IsSpace)

	// TODO: refactor this
	if strings.HasPrefix(input, "!") {
		n = true
		// Remove !
		input = strings.Split(input, "!")[1]
		// remove spaces
		// input = strings.TrimLeftFunc(input, unicode.IsSpace)
		input = strings.TrimFunc(input, unicode.IsSpace)
	}
	var p []string = strings.Split(input, " ")

	// ToDo: error handling. check if len(p) > 1
	input = strings.Join(p[1:], " ")

	sm := searchMatch{
		value:    input,
		negative: n,
	}

	return p[0], sm
}

func parseInput(input string) (SearchParams, error) {

	var parseAnd []string
	var sp SearchParams = SearchParams{}

	parseAnd = strings.Split(input, "&")

	for _, v := range parseAnd {

		k, match := parseNot(v)

		switch k {
		case "ReqMethod":
			sp.ReqMethod = append(sp.ReqMethod, match)
		case "ReqHeader":
			sp.ReqHeader = append(sp.ReqHeader, match)
		case "ReqContentType":
			sp.ReqContentType = append(sp.ReqContentType, match)
		case "ReqBody":
			sp.ReqBody = append(sp.ReqBody, match)
		case "ResHeader":
			sp.ResHeader = append(sp.ResHeader, match)
		case "ResContentType":
			sp.ResContentType = append(sp.ResContentType, match)
		case "ResBody":
			sp.ResBody = append(sp.ResBody, match)
		default:
			return SearchParams{}, fmt.Errorf("unsupported search field: %s. Use help (in menu) for list of supported field.", k)
		}

	}

	return sp, nil
}

func (s *Search) parseSearch(input string) error {
	// todo: error handler
	re, err := parseInput(input)
	s.Regexp = re

	return err

}
