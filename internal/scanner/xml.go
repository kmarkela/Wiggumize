package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"regexp"
)

func buidXMLCheck() Check {

	return Check{
		Description: "This module is searching for XML in request parameters",
		Execute:     searchForXML,
		Config:      nil,
	}
}

func searchForXML(p parser.HistoryItem, c *Check) []Finding {

	// if p.Method == "POST" {
	// 	fmt.Println(p.Params)
	// }

	rePatern := `\<.*\>`

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
		Description: "possible XML in a parameter",
		Evidens:     p.Params,
		Details:     "URL:" + p.URL,
	}
	findings = append(findings, finding)
	return findings
}
