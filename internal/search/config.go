package search

import (
	"Wiggumize/cli"
	"strconv"
)

type SearchConfig struct {
	Output          outputType
	CaseInsensitive bool
}

type outputType string

const (
	endpoint outputType = "endpoint"
	headers  outputType = "headers"
	reqOnly  outputType = "reqOnly"
	full     outputType = "full"
)

func (s *Search) togleCase() {

	var ops []string

	ops = append(ops, "true")
	ops = append(ops, "false")
	caseInsensitive := cli.GetSelect("caseInsensitive", ops, strconv.FormatBool(s.Config.CaseInsensitive))

	// TODO: error handling
	s.Config.CaseInsensitive, _ = strconv.ParseBool(caseInsensitive)

}

func (s *Search) togleOutput() {

	var ops []string

	ops = append(ops, "endpoint")
	ops = append(ops, "headers")
	ops = append(ops, "reqOnly")
	ops = append(ops, "full")
	output := cli.GetSelect("Output", ops, string(s.Config.Output))

	s.Config.Output = outputType(output)

}

func (s *Search) handleConfig() {

	var ops []string

	ops = append(ops, "Output")
	ops = append(ops, "CaseInsensitive")
	ops = append(ops, "Back")
	subMenu := cli.GetSelect("Config", ops, "Output")

	switch subMenu {
	case "Output":
		s.togleOutput()
	case "CaseInsensitive":
		s.togleCase()
	case "Back":
		return
	}

	s.handleConfig()
}
