package cli

import (
	"github.com/fatih/color"
)

// TODo: refactor: colors

func header() {

	nameASCIIFirst := `
	__       __  __                                              __                     
	|  \  _  |  \|  \                                            |  \                    
	| $$ / \ | $$ \$$  ______    ______   __    __  ______ ____   \$$ ________   ______  
	| $$/  $\| $$|  \ /      \  /      \ |  \  |  \|      \    \ |  \|        \ /      \ 
	| $$  $$$\ $$| $$|  $$$$$$\|  $$$$$$\| $$  | $$| $$$$$$\$$$$\| $$ \$$$$$$$$|  $$$$$$\
	| $$ $$\$$\$$| $$| $$  | $$| $$  | $$| $$  | $$| $$ | $$ | $$| $$  /    $$ | $$    $$
	| $$$$  \$$$$| $$| $$__| $$| $$__| $$| $$__/ $$| $$ | $$ | $$| $$ /  $$$$_ | $$$$$$$$
	| $$$    \$$$| $$ \$$    $$ \$$    $$ \$$    $$| $$ | $$ | $$| $$|  $$    \ \$$     \
	 \$$      \$$ \$$ _\$$$$$$$ _\$$$$$$$  \$$$$$$  \$$  \$$  \$$ \$$ \$$$$$$$$  \$$$$$$$
			 |  \__| $$|  \__| $$   `

	nameASCIIDescription := "Web Traffic 4nalizer"

	nameASCIILast := `
			  \$$    $$ \$$    $$
			   \$$$$$$   \$$$$$$
	`

	// Define the colors
	red := color.New(color.FgBlue)
	boldYellow := color.New(color.FgYellow, color.Bold).Add(color.Underline)

	red.Print(nameASCIIFirst)
	boldYellow.Printf(nameASCIIDescription)
	red.Println(nameASCIILast)

}

func Greet() {
	header()

}
