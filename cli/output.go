package cli

import (
	scan "Wiggumize/internal/scanner"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildACheckMD(key string, val []scan.Finding) string {

	var content string
	for i, f := range val {
		content += "### Finding " + strconv.Itoa(i) + ". - " + f.Description + "\n"
		content += "__Host: " + f.Host + "__ \n\n"
		content += "_Evidens:_\n\n```\n" + f.Evidens + "\n```\n"
		if f.Details != "" {
			content += "_More Details:_\n\n```\n" + f.Details + "\n```\n"
		}
	}

	return content

}

func OutputToMD(scanner *scan.Scanner, scope []string, filename string) error {

	content := "# Wiggumize Report\n\n"
	content += "__Scope:__\n"

	for _, host := range scope {
		content += "- " + host + "\n"
	}
	content += "\n\n"
	content += "__List of Checks:__\n"

	for key, val := range scanner.ChecksMap {
		content += "- __" + key + ":__ " + val.Description + "\n"
	}

	content += strings.Repeat("-", 20)
	content += "\n\n"

	for key, val := range scanner.Results {

		// TODO: move epty checrs to the end or remove them
		if len(val) == 0 {
			continue
		}

		content += "## " + key + "\n"
		content += "> " + scanner.ChecksMap[key].Description + "\n"
		content += buildACheckMD(key, val)

	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	file.WriteString(content)

	return nil

}
