package cli

import (
	"github.com/AlecAivazis/survey/v2"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opts,
		PageSize: 15,
	}
	survey.AskOne(prompt, &res)

	return res
}
