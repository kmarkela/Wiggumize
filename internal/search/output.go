package search

import (
	"Wiggumize/cli"
	parser "Wiggumize/internal/trafficParser"
	"strings"
)

func printEndpoints(f []parser.HistoryItem, cPrinter *cli.ColorPrinter) {
	for i, v := range f {
		cPrinter.Blue.Printf(" Match # %d: ", i+1)
		cPrinter.White.Printf("%s\n", v.URL)
		// fmt.Printf("# Match # %d: %s \n", i+1, v.URL)
	}

}

func printHeaders(f []parser.HistoryItem, cPrinter *cli.ColorPrinter, delimiter string) {
	for i, v := range f {
		cPrinter.Blue.Printf("# Match %d. \n", i+1)
		cPrinter.Blue.Printf("## Endpoint: ")
		cPrinter.White.Printf("%s\n", v.URL)
		cPrinter.Blue.Printf("## ReqHeaders: \n")
		cPrinter.White.Printf("%s\n\n", v.ReqHeaders)
		cPrinter.Blue.Printf("## ResHeaders: \n")
		cPrinter.White.Printf("%s\n\n", v.ResHeaders)
		cPrinter.Yellow.Printf("%s\n\n", delimiter)
	}
}

func printReq(f []parser.HistoryItem, cPrinter *cli.ColorPrinter, delimiter string) {
	for i, v := range f {
		cPrinter.Blue.Printf("# Match %d. \n", i+1)
		cPrinter.Blue.Printf("## Endpoint: ")
		cPrinter.White.Printf("%s\n", v.URL)
		cPrinter.Blue.Printf("## Headers: \n")
		cPrinter.White.Printf("%s\n\n", v.ReqHeaders)
		cPrinter.Blue.Printf("## Body: \n")
		cPrinter.White.Printf("%s\n\n", v.ReqBody)
		cPrinter.Yellow.Printf("%s\n\n", delimiter)
	}
}

// TODO: refactor code repetition
func printFull(f []parser.HistoryItem, cPrinter *cli.ColorPrinter, delimiter string) {
	for i, v := range f {

		cPrinter.Blue.Printf("# Match %d. \n", i+1)
		cPrinter.Blue.Printf("## Endpoint: ")
		cPrinter.White.Printf("%s\n", v.URL)
		cPrinter.Blue.Printf("## ReqHeaders: \n")
		cPrinter.White.Printf("%s\n\n", v.ReqHeaders)
		cPrinter.Blue.Printf("## ReqBody: \n")
		cPrinter.White.Printf("%s\n\n", v.ReqBody)
		cPrinter.Blue.Printf("## ResHeaders: \n")
		cPrinter.White.Printf("%s\n\n", v.ResHeaders)
		cPrinter.Blue.Printf("## ResBody: \n")
		cPrinter.White.Printf("%s\n\n", v.ResBody)
		cPrinter.Yellow.Printf("%s\n\n", delimiter)
	}
}

func (s *Search) output() {

	var delimiter string = strings.Repeat("$", 64)
	cPrinter := cli.NewColorPrinter()

	cPrinter.AddAttributeString(cPrinter.Green, "Bold")
	cPrinter.AddAttributeString(cPrinter.Cyan, "Underline")
	cPrinter.Green.Printf("Found %d matches.\n", len(s.Found))
	// cPrinter.Cyan.Printf(delimiter)

	if len(s.Found) == 0 {
		return
	}

	switch s.Config.Output {
	case "endpoint":
		printEndpoints(s.Found, cPrinter)
	case "headers":
		printHeaders(s.Found, cPrinter, delimiter)
	case "reqOnly":
		printReq(s.Found, cPrinter, delimiter)
	case "full":
		printFull(s.Found, cPrinter, delimiter)
	}
}
