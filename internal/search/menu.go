package search

import (
	"Wiggumize/cli"
	"fmt"
)

func handleMenu() {

	input := printMenu()

	switch input {
	case "Help":
		printHelpMsg()
	case "Config":
		// TODO: Config
		fmt.Println("doNOtImplemented")
	default:
		return
	}
}

func printMenu() string {
	menuOptions := []string{"Help", "Config", "Back"}

	return cli.GetSelect("Choose an option: ", menuOptions)
}
