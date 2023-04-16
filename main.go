package main

import (
	"fmt"

	"Wiggumize/cli"
	scan "Wiggumize/internal/scanner"
	parser "Wiggumize/internal/trafficParser"
	"Wiggumize/utils"
)

func main() {

	// Get Cli Parameters
	var params *cli.Parameters
	params = &cli.Parameters{}
	params.Parse()

	// Print Greatings
	cli.Greet()

	// ############################
	// Prepare History to work with
	// ############################
	var browseHistory *parser.BrowseHistory
	browseHistory = &parser.BrowseHistory{
		RequestsList: []parser.HistoryItem{},
		ListOfHosts:  utils.Set{},
	}

	// futureprofing. will be in use in v2 .
	switch {
	case params.FilePath != "":
		XMLParser := parser.XMLParser{}
		err := XMLParser.PopulateHistory(params.FilePath, browseHistory)
		if err != nil {
			//TODO: add proper logging
			fmt.Println(err)
			return
		}
	// case params.Proxy
	// case params.API
	default:
		// Should never happen
		params.Usage()
		panic("Can't parse CLI parameters")
	}

	// ############################
	// Analyse History
	// ############################

	// filter scope
	scopeHosts := cli.Checkboxes("Choose hosts in Scope:", browseHistory.ListOfHosts.Keys())
	browseHistory.FilterByHost(scopeHosts)

	scanner, err := scan.SannerBuilder()
	if err != nil {
		fmt.Println("Cannot Start Scanner!", err)
		return
	}

	scanner.RunAllChecks(browseHistory)

	for key, val := range scanner.ChecksMap {
		// TODO: move it to CLI
		fmt.Println("#################################")
		fmt.Printf("####### %s ################\n", key)
		fmt.Println("#################################")
		fmt.Println(val.Description)
		fmt.Println("--------------------------------")
		fmt.Printf("Found: %d\n", len(val.Results))
		fmt.Println("--------------------------------")

		for i, m := range val.Results {
			fmt.Println("--------------------------------")
			fmt.Printf("Finding #%d\n", i)
			fmt.Printf("Description: %s\n", m.Description)
			fmt.Printf("Evidens: %s\n", m.Evidens)
			fmt.Printf("Host: %s\n", m.Host)
			fmt.Printf("Details: %s\n", m.Details)
		}

	}

}
