package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

type service struct {
	Name            string
	NotFoundMessage string
}

type notFoundConfig struct {
	Title    string
	Services []service
}

func buidNotFoundCheck() Check {
	nc, err := get404Messages()
	if err != nil {
		fmt.Println("Cennot get rules for Secrets")
	}

	return Check{
		Description: "This module is searching for secrets (eg. API keys) in Req\\Res",
		Execute:     searchForNotFound,
		Config:      nc,
	}
}

func get404Messages() (notFoundConfig, error) {
	// Open the TOML file
	// TODO: use const from config/consts.go
	// TODO: Refactor. Code repetition on each TOML read
	file, err := os.Open("internal/config/scan/404sub.toml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return notFoundConfig{}, err
	}
	defer file.Close()

	var nc notFoundConfig
	err = toml.NewDecoder(file).Decode(&nc)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return notFoundConfig{}, err
	}
	return nc, nil

}

func searchForNotFound(p parser.HistoryItem, c *Check) []Finding {

	var findings []Finding

	if p.Status != "404" {
		return findings
	}

	for _, message := range c.Config.(notFoundConfig).Services {
		if !strings.Contains(p.Response, message.NotFoundMessage) {
			continue
		}

		finding := Finding{Host: p.Host,
			Description: "404 from hosting service",
			Evidens:     message.NotFoundMessage,
			Details:     "Service: " + message.Name + "URL:" + p.URL,
		}
		findings = append(findings, finding)
		break

	}
	return findings
}
