package utils

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

func ParseTOML(filename string, obj *interface{}) error {
	// Open the TOML file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(&obj)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return err
	}

	return nil

}
