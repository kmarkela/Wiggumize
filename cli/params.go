package cli

import (
	"flag"
	"fmt"
	"os"
)

type Parameters struct {
	FilePath string
	Output   string
}

func (p *Parameters) Parse() {
	flag.StringVar(&p.FilePath, "f", "", "path to XML file with burp history")
	flag.StringVar(&p.Output, "o", "retport.md", "path to output")
	flag.Parse()

	// Check if the file path flag is set.
	if p.FilePath == "" {
		fmt.Fprintf(os.Stderr, "Error: missing file path parameter\n")
		flag.Usage()
		os.Exit(1)
	}

}

func (p *Parameters) Usage() {
	// Expose Usage
	flag.Usage()
}
