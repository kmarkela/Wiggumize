package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"regexp"
)

func buidSSRFCheck() Check {
	return Check{
		Description: "This module is searching for URL in request parameters.",
		Execute:     searchForURLs,
		Config:      nil,
	}
}

func searchForURLs(p parser.HistoryItem, c *Check) []Finding {

	rePatern := `(https?):\/\/[^\s\/$.?#].[^\s\/]*\/?`

	regex, err := regexp.Compile(rePatern)
	if err != nil {
		fmt.Printf("Error compiling regex pattern: %s\n", err)
	}

	match := regex.FindString(p.Params)

	if match == "" {
		return nil
	}

	var findings []Finding

	finding := Finding{Host: p.Host,
		Description: "URL in a parameter",
		Evidens:     p.Params,
		Details:     "URL:" + p.URL,
	}
	findings = append(findings, finding)
	return findings
}
