package search

import (
	"Wiggumize/utils"
	"fmt"
)

func printHelpMsg() {
	help := "\nRegexp Search: \n\n"
	help += "Avaliable search fields: \n"

	searchFields := utils.GetStructFieldNames(SearchParams{})

	for _, name := range searchFields {
		help += "- " + name + "\n"
	}

	help += "\nAvaliable search operators: \n"

	help += "- & - AND\n"
	help += "- ! - NOT\n\n"

	help += "Search Example: \n"
	help += "ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success\n"

	fmt.Println(help)
}
