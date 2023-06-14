package search

import (
	"Wiggumize/cli"
)

func handleMenu(s *Search) {

	input := printMenu()

	switch input {
	case "Help":
		printHelpMsg()
	case "Config":
		s.handleConfig()
	default:
		return
	}
}

func printMenu() string {
	menuOptions := []string{"Help", "Config", "Back"}

	return cli.GetSelect("Choose an option: ", menuOptions, "Help")
}
