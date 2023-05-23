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

type secretsConfig struct {
	Title string
	Rules []rule
}

type SecretMatch struct {
	MatchingString string
	Description    string
}

// var secret = Check{
// 	Description: "This module is searching for secrets (eg. API keys) in Req\\Res",
// 	Execute:     searchForSecrets,
// 	CheckReq:    true,
// 	CheckRes:    true,
// 	config: secretsConfig{
// 		Title: "",

// 	}
// }

func buidSecretCheck() Check {
	sc, err := getRules()
	if err != nil {
		fmt.Println("Cennot get rules for Secrets")
	}

	return Check{
		Description: "This module is searching for secrets (eg. API keys)",
		Execute:     searchForSecrets,
		Config:      sc,
	}
}

func getRules() (secretsConfig, error) {
	// Open the TOML file
	// TODO: use const from config/consts.go
	// TODO: Refactor. Code repetition on each TOML read
	file, err := os.Open("internal/config/scan/secrets.toml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return secretsConfig{}, err
	}
	defer file.Close()

	var sc secretsConfig
	err = toml.NewDecoder(file).Decode(&sc)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return secretsConfig{}, err
	}
	return sc, nil

}

func checkString(s string, sc secretsConfig) []SecretMatch {

	// TODO: change to set
	var matchList []SecretMatch
	for _, rule := range sc.Rules {

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

func buidFindingsFromSecret(ml []SecretMatch, direction string, host string) []Finding {

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

func searchForSecrets(p parser.HistoryItem, c *Check) []Finding {

	// check req/res
	listOfMatches := checkString(p.Request, c.Config.(secretsConfig))
	findings := buidFindingsFromSecret(listOfMatches, "Req", p.Host)

	listOfMatches = checkString(p.Response, c.Config.(secretsConfig))
	findings = append(findings, buidFindingsFromSecret(listOfMatches, "Res", p.Host)...)

	return findings
}
