package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"os"

	"regexp"

	"github.com/pelletier/go-toml"
)

type rule struct {
	Description string
	ID          string
	Regex       string
}

type config struct {
	Title string
	Rules []rule
}

type SecretMatch struct {
	MatchingString string
	Description    string
}

var secret = Check{
	Description: "This module is searching for secrets (eg. API keys) in Req\\Res",
	Execute:     searchForSecrets,
	CheckReq:    true,
	CheckRes:    true,
}

func getRules() ([]rule, error) {
	// Open the TOML file
	// TODO: use const from config/consts.go
	// TODO: Refactor. Code repetition on each TOML read
	file, err := os.Open("internal/config/scan/secrets.toml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var config config
	err = toml.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return nil, err
	}

	return config.Rules, nil

}

func checkString(s string) []SecretMatch {

	var rules []rule
	// TODO: hendle errors
	rules, _ = getRules()

	// TODO: change to set
	var matchList []SecretMatch

	for _, rule := range rules {

		regex, err := regexp.Compile(rule.Regex)
		if err != nil {
			fmt.Printf("Error compiling regex pattern: %s\n", err)
			continue
		}

		match := regex.FindString(s)

		if match != "" {
			matchList = append(matchList, SecretMatch{MatchingString: match, Description: rule.Description})
			// fmt.Println(match)
		}

	}

	return matchList
}

func buidFindings(ml []SecretMatch, direction string, host string) []Finding {

	findings := []Finding{}
	for _, item := range ml {
		finding := Finding{
			Host:        host,
			Description: item.Description,
			Evidens:     item.MatchingString,
			Details:     direction,
		}
		findings = append(findings, finding)
	}

	return findings
}

func searchForSecrets(p parser.HistoryItem) []Finding {

	// check req/res
	listOfMatches := checkString(p.Request)
	findings := buidFindings(listOfMatches, "Req", p.Host)

	listOfMatches = checkString(p.Response)
	findings = append(findings, buidFindings(listOfMatches, "Res", p.Host)...)

	return findings
}
