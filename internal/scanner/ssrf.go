package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"regexp"
)

func buidSSRFCheck() Check {
	ec, err := getExt()
	if err != nil {
		fmt.Println("Cennot get rules for LFI")
	}

	return Check{
		Description: "This module is searching for filenames in request parameters",
		Execute:     searchForURLs,
		Config:      ec,
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
