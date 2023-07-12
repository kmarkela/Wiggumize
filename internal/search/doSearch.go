package search

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"regexp"
	"sync"
)

func (s *Search) waitForResults() {
	for {
		select {
		case msg := <-s.channel: // recived message
			s.Found = append(s.Found, msg)
		default:
			// fmt.Println("asdsadasd")
		}

	}
}

func (s *Search) doSearch(input string) {

	// TODO: refactor with return instead for pointers
	err := s.parseSearch(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	s.channel = make(chan parser.HistoryItem)

	go s.waitForResults()

	for _, item := range s.PHistory.RequestsList {

		wg.Add(1) // add a worker to the waitgroup
		go s.checkRegexp(item, &wg)

	}

	wg.Wait()

}

func (s *Search) checkRegexp(p parser.HistoryItem, wg *sync.WaitGroup) {
	defer wg.Done() // signal that the worker has finished

	if !regexMatch(s.Regexp.ReqMethod, p.Method, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ReqHeader, p.ReqHeaders, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ReqContentType, p.ReqContentType, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ReqBody, p.ReqBody, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ResHeader, p.ResHeaders, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ResContentType, p.ResContentType, s.Config.CaseInsensitive) {
		return
	}

	if !regexMatch(s.Regexp.ResBody, p.ResBody, s.Config.CaseInsensitive) {
		return
	}
	s.channel <- p
}

func regexMatch(m []searchMatch, st string, ci bool) bool {

	// Case insesitive
	var prefix string = ""
	if ci {
		prefix = "(?i)"
	}

	// No search rexexp
	if len(m) == 0 {
		return true
	}

	//empty Body
	if len(st) == 0 {
		return false
	}

	for _, v := range m {

		match, _ := regexp.MatchString(prefix+v.value, st)

		// fmt.Printf("Neg - %t, val - %s; Match - %t; String: \n%s \n\n\n", v.negative, v.value, match, st)

		// Found Negative or doesn't find positive
		if (match && v.negative) || (!match && !v.negative) {
			return false
		}
	}

	return true
}
