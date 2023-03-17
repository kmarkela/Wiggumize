package cli

import (
	"flag"
	"fmt"
	"os"
)

type Parameters struct {
	FilePath string
}

func (p *Parameters) Parse() {
	flag.StringVar(&p.FilePath, "file", "", "path to XML file")
	flag.Parse()

	// Check if the file path flag is set.
	if p.FilePath == "" {
		fmt.Fprintf(os.Stderr, "Error: missing file path parameter\n")
		flag.Usage()
		os.Exit(1)
	}
}
