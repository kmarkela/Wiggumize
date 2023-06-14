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

	help += "- contain\n"
	help += "- eq\n"
	help += "- notEq\n"
	help += "- notContain\n\n"

	help += "Search Example: \n"
	help += "ReqMethod.eq POST & ReqBody.contain admin & ResContentType.notContain HTML & ResBody.contain success\n"

	fmt.Println(help)
}
