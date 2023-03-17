package main

import (
	"fmt"
	"os"

	"Wiggumize/cli"
	"Wiggumize/internal/parser"
	"Wiggumize/internal/proxy"
	// "Wiggumize/internal/passive"
)

func main() {

	// Get Cli Parameters
	var params *cli.Parameters
	params = &cli.Parameters{}
	params.Parse()

	// Print Greatings
	cli.Greet()

	// TODO: normalisations. shoudl be watever XML/API
	// Filter XML
	var filteredXML *parser.XMLParser
	filteredXML = xmlProcess(params)
	fmt.Println(len(filteredXML.ItemElements))

	// test := "test"
	// m, b := passive.Find(test)
	// fmt.Println(b)
	// fmt.Println(m)

	// for _, item := range filteredXML.ItemElements {

	// 	// if bool(item.Request.Base64) {
	// 	// 	// item.Response.Value = parser.B64Decode(item.Response.Value)
	// 	// 	fmt.Println(parser.B64Decode(item.Response.Value))
	// 	// }

	// 	// fmt.Println(item.Response.Value)
	// 	s := parser.B64Decode(item.Response.Value)

	// 	matched, match := passive.Find(s)
	// 	//TODO: refactor this
	// 	if matched {
	// 		for _, m := range match {
	// 			fmt.Println("======FOUND SECRET========================")
	// 			fmt.Printf("Description: %s\n", m.Description)
	// 			fmt.Print("Value: ")
	// 			fmt.Println(m.MatchingString)
	// 			fmt.Printf("URL: %s\n", item.URL)
	// 			// fmt.Println("WHole Responce:")
	// 			// fmt.Println(s)
	// 			fmt.Println("======================================")
	// 		}
	// 	}

	// }

	var proxyIntance *proxy.Proxy
	proxyIntance = &proxy.Proxy{}
	proxyIntance.Start(1337)

}

func xmlProcess(params *cli.Parameters) *parser.XMLParser {
	// Parse XML
	xmlParser := &parser.XMLParser{}
	err := xmlParser.Parse(params.FilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML file: %v\n", err)
		os.Exit(1)
	}

	// Get List of Hosts from history
	listOfHosts := xmlParser.ListOfHosts()

	selectedHosts := cli.Checkboxes("Choose hosts os your interest:", listOfHosts)

	// Filter history to get only req/res for selected hosts
	xmlParser.FilterByHost(selectedHosts)

	return xmlParser
}
