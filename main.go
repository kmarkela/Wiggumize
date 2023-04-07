package main

import (
	"fmt"

	"Wiggumize/cli"
	"Wiggumize/internal/parser"
	"Wiggumize/internal/passive"
	"Wiggumize/utils"
)

func main() {

	// Get Cli Parameters
	var params *cli.Parameters
	params = &cli.Parameters{}
	params.Parse()

	// Print Greatings
	cli.Greet()

	// ##########
	// Parse Data
	// ##########
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

	selectedHosts := cli.Checkboxes("Choose hosts os your interest:", browseHistory.ListOfHosts.Keys())
	browseHistory.FilterByHost(selectedHosts)

	for _, item := range browseHistory.RequestsList {

		// if bool(item.Request.Base64) {
		// 	// item.Response.Value = parser.B64Decode(item.Response.Value)
		// 	fmt.Println(parser.B64Decode(item.Response.Value))
		// }

		// fmt.Println(item.Response.Value)
		matched, match := passive.Find(item.Response)
		//TODO: refactor this
		if matched {
			for _, m := range match {
				fmt.Println("======FOUND SECRET========================")
				fmt.Printf("Description: %s\n", m.Description)
				fmt.Print("Value: ")
				fmt.Println(m.MatchingString)
				fmt.Printf("URL: %s\n", item.URL)
				// fmt.Println("WHole Responce:")
				// fmt.Println(s)
				fmt.Println("======================================")
			}
		}

	}

}

// func (p *XMLParser) xmlProcess(params *cli.Parameters) *parser.XMLParser {
// 	// Parse XML
// 	xmlParser := &parser.XMLParser{}
// 	err := xmlParser.Parse(params.FilePath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error parsing XML file: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Get List of Hosts from history
// 	listOfHosts := xmlParser.ListOfHosts()

//

// 	// Filter history to get only req/res for selected hosts
// 	xmlParser.FilterByHost(selectedHosts)

// 	return xmlParser
// }
