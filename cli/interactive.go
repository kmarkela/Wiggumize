package cli

import (
	"Wiggumize/utils"

	"github.com/AlecAivazis/survey/v2"
)

func Checkboxes(label string, opts []string) utils.Set {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opts,
		PageSize: 15,
	}
	survey.AskOne(prompt, &res)

	hostSet := utils.Set{}
	for _, item := range res {
		hostSet.Add(item)
	}

	return hostSet
}
