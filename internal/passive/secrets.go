package passive

import (
	"fmt"
	"os"

	"regexp"

	"github.com/pelletier/go-toml"
)

type Rule struct {
	Description string
	ID          string
	Regex       string
}

type Config struct {
	Title string
	Rules []Rule
}

type SecretMatch struct {
	MatchingString string
	Description    string
}

func getRules() ([]Rule, error) {
	// Open the TOML file
	file, err := os.Open("internal/config/secrets.toml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var config Config
	err = toml.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return nil, err
	}

	return config.Rules, nil

}

func Find(s string) (bool, []SecretMatch) {

	var rules []Rule
	// TODO: hendle errors
	rules, _ = getRules()

	var matchList []SecretMatch

	for _, rule := range rules {

		regex, err := regexp.Compile(rule.Regex)
		if err != nil {
			fmt.Printf("Error compiling regex pattern: %s\n", err)
			continue
		}

		match := regex.FindString(s)

		if match != "" {
			matchList = append(matchList, SecretMatch{MatchingString: match, Description: rule.Description})
			// fmt.Println(match)
		}

	}

	return len(matchList) > 0, matchList
}
