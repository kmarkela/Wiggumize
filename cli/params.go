package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Parameters struct {
	FilePath string
	Output   string
	Action   string
}

func (p *Parameters) Parse() {
	flag.StringVar(&p.FilePath, "f", "", "path to XML file with burp history")
	flag.StringVar(&p.Output, "o", "retport.md", "path to output")
	flag.StringVar(&p.Action, "a", "scan", "Action. scan/search")
	flag.Parse()

	// Check if the file path flag is set.
	if p.FilePath == "" {
		fmt.Fprintf(os.Stderr, "Error: missing file path parameter\n")
		flag.Usage()
		os.Exit(1)
	}

	// make it case insesitive
	p.Action = strings.ToLower(p.Action)

}

func (p *Parameters) Usage() {
	// Expose Usage
	flag.Usage()
}
