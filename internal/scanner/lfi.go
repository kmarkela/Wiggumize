package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"fmt"
	"os"
	"regexp"

	"github.com/pelletier/go-toml"
)

type Extension struct {
	Description string
	Ext         string
}

type ExtsConfig struct {
	Title      string
	Extensions []Extension
}

func getExt() (ExtsConfig, error) {
	// Open the TOML file
	// TODO: use const from config/consts.go
	// TODO: Refactor. Code repetition on each TOML read
	file, err := os.Open("internal/config/scan/lfi.toml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ExtsConfig{}, err
	}
	defer file.Close()

	var ec ExtsConfig
	err = toml.NewDecoder(file).Decode(&ec)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return ExtsConfig{}, err
	}
	return ec, nil

}

func buidLfiCheck() Check {
	ec, err := getExt()
	if err != nil {
		fmt.Println("Cennot get rules for LFI")
	}

	return Check{
		Description: "This module is searching for filenames in request parameters. Could be an indication of possible LFI",
		Execute:     searchForFiles,
		Config:      ec,
	}
}

func searchForFiles(p parser.HistoryItem, c *Check) []Finding {

	rePatern := ".*\\.("

	// Build a pattern for filename
	for i, ext := range c.Config.(ExtsConfig).Extensions {
		//`.*\.(txt|php|exe)$`
		rePatern = rePatern + ext.Ext[1:] // add exts w\o leading dot
		if i == len(c.Config.(ExtsConfig).Extensions)-1 {
			rePatern = rePatern + ").*"
			break
		}
		rePatern = rePatern + "|"

	}
	// fmt.Println(rePatern)

	// check req/res

	regex, err := regexp.Compile(rePatern)
	if err != nil {
		fmt.Printf("Error compiling regex pattern: %s\n", err)
	}

	match := regex.FindString(p.Params)

	if match == "" {
		return nil
	}

	var findings []Finding

	finding := Finding{Host: p.Host,
		Description: "filename in a parameter",
		Evidens:     p.Params,
		Details:     "URL: " + p.URL,
	}
	findings = append(findings, finding)
	return findings
}
