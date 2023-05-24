package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"strings"
)

type ParamsMap struct {
	Name        string
	Description string
	Hosts       map[string]ParsedParams
}

type ParsedParams struct {
	Endpoints map[string]EndpointParams
}

type EndpointParams struct {
	Params map[string]string
}

func parseGETParams(params string) EndpointParams {
	paramsList := strings.Split(params, "&")

	e := EndpointParams{
		Params: map[string]string{},
	}
	for _, p := range paramsList {
		param := strings.Split(p, "=")
		e.Params[param[0]] = param[1]
	}

	return e
}

func parseParams(p parser.HistoryItem) (string, string, EndpointParams) {

	if p.Params == "" {
		return "", "", EndpointParams{}
	}

	if p.Method == "GET" {
		return p.Host, p.Path, parseGETParams(p.Params)
	}

	return "", "", EndpointParams{}

}
