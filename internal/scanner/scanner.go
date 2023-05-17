package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"sync"
)

type Scanner struct {
	ChecksMap map[string]Check
	channel   chan channalMessage
	Results   map[string][]Finding
}

type Check struct {
	Description string
	Execute     func(parser.HistoryItem, *Check) []Finding
	Executed    bool
	Config      interface{}
}

type Finding struct {
	Host        string
	Description string
	Evidens     string
	Details     string //TODO: add just Json implementation
}

type Result struct {
	CheckName string
	Findings  []Finding
}

type channalMessage struct {
	checkName string
	Findings  []Finding
}

func SannerBuilder() (*Scanner, error) {
	/*
		1. Creates instance of Scanner
		2. populates ChecksMap
	*/

	// Cehck defined in the cehck's file
	var checksMap = map[string]Check{
		"Secrets": buidSecretCheck(),
		"lfi":     buidLfiCheck(),
		"ssrf":    buidSSRFCheck(),
	}

	scanner := &Scanner{
		ChecksMap: checksMap,
		channel:   make(chan channalMessage, 128),
		Results:   make(map[string][]Finding),
	}

	// TODO: get list of scans from config
	return scanner, nil

}

func (s *Scanner) runChecks(r parser.HistoryItem, wg *sync.WaitGroup) {

	defer wg.Done() // signal that the worker has finished

	for key, check := range s.ChecksMap {
		findings := check.Execute(r, &check)

		s.channel <- channalMessage{
			checkName: key,
			Findings:  findings,
		}
		// check.Results = append(s.ChecksMap[key].Results, results...)
		// s.ChecksMap[key] = check

	}
}

func (s *Scanner) waitForResults() {
	for {
		select {
		case msg := <-s.channel: // recived message
			s.Results[msg.checkName] = append(s.Results[msg.checkName], msg.Findings...)
		default:
		}

	}
}

func (s *Scanner) RunAllChecks(b *parser.BrowseHistory) {

	var wg sync.WaitGroup

	go s.waitForResults()

	for _, item := range b.RequestsList {

		wg.Add(1) // add a worker to the waitgroup
		go s.runChecks(item, &wg)

	}
	wg.Wait()

}
