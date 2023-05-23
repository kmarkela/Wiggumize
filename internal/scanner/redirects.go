package scan

import (
	parser "Wiggumize/internal/trafficParser"
)

func buidRedirectCheck() Check {
	return Check{
		Description: "This module is searching for Redirects.",
		Execute:     searchForRedirects,
		Config:      nil,
	}
}

func searchForRedirects(p parser.HistoryItem, c *Check) []Finding {

	var findings []Finding
	var finding Finding

	var codeMap map[string]string = make(map[string]string)

	codeMap["301"] = "Moved Permanently"
	codeMap["302"] = "Found"
	codeMap["303"] = "See Other"
	codeMap["307"] = "Temporary Redirect"
	codeMap["308"] = "Permanent Redirect"

	// if !strings.HasPrefix(p.Status, "3") {
	// 	return findings
	// }

	if _, ok := codeMap[p.Status]; !ok {
		return findings
	}

	if p.Params != "" {
		finding = Finding{Host: p.Host,
			Description: "Redirect Found",
			Evidens:     p.Response,
			Details:     "Req Parameters:" + p.Params,
		}
	} else {
		finding = Finding{Host: p.Host,
			Description: "Redirect Found",
			Evidens:     p.Response,
		}
	}

	findings = append(findings, finding)
	return findings
}
