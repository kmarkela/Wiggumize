package scan

import (
	parser "Wiggumize/internal/trafficParser"
)

type Scanner struct {
	ChecksMap map[string]Check
}

type Check struct {
	Description string
	Execute     func(parser.HistoryItem) []Finding
	Executed    bool
	CheckReq    bool // if true request will be checked
	CheckRes    bool // if true responce will be checked
	Results     []Finding
}

type Finding struct {
	Host        string
	Description string
	Evidens     string
	Details     string //TODO: add just Json implementation
}

func SannerBuilder() (*Scanner, error) {
	/*
		1. Creates instance of Scanner
		2. populates ChecksMap
	*/

	// Cehck defined in the cehck's file
	var checksMap = map[string]Check{
		"Secrets": secret,
	}

	scanner := &Scanner{}
	scanner.ChecksMap = checksMap

	// TODO: get list of scans from config
	return scanner, nil

}

func (s *Scanner) RunACheck(b *parser.BrowseHistory, checkName string) {

	// Greb the check from the Map
	theCheck := s.ChecksMap[checkName]

	for _, item := range b.RequestsList {
		// execute check on each
		// ToDo: do via rutine
		theCheck.Results = append(theCheck.Results, theCheck.Execute(item)...)
	}

	s.ChecksMap[checkName] = theCheck

}

func (s *Scanner) RunAllChecks(b *parser.BrowseHistory) {

	// Run all checks in separate rutine
	for key, _ := range s.ChecksMap {
		go s.RunACheck(b, key)
	}

}
