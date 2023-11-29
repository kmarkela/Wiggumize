package scan

import (
	parser "Wiggumize/internal/trafficParser"
	"encoding/json"
	"fmt"
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
	Params Param
	Method string
}

type Param map[string]string

func parseGETParams(params string) Param {
	paramsList := strings.Split(params, "&")

	returnParams := Param{}

	for _, p := range paramsList {
		param := strings.Split(p, "=")
		if len(param) < 2 {
			continue
		}
		returnParams[param[0]] = param[1]
	}

	return returnParams
}

func parsePOSTParams(params string, ct string) Param {
	// e := EndpointParams{
	// 	Params: map[string]string{},
	// }

	returnParams := Param{}

	switch ct {
	case "application/x-www-form-urlencoded":
		returnParams = parseGETParams(params)
	case "application/json":
		parseJSONBody(params, returnParams)
	// case "multipart/form-data":
	// 	e = parseMultipartFormBody(params)
	default:
		returnParams["Unable to parse content type"] = ct
	}

	return returnParams
}

func parseJSONBody(params string, returnParams Param) {

	var jsonObj interface{}
	err := json.Unmarshal([]byte(params), &jsonObj)
	if err != nil {
		// return err
		// TODO: add errorHandling
	}

	parseJSONValue(jsonObj, "", returnParams)

}

func parseJSONValue(value interface{}, prefix string, returnParams Param) {
	switch value := value.(type) {
	case map[string]interface{}:
		for key, subvalue := range value {
			subprefix := fmt.Sprintf("%s.%s", prefix, key)
			if prefix == "" {
				subprefix = key
			}
			parseJSONValue(subvalue, subprefix, returnParams)
		}
	case []interface{}:
		for i, subvalue := range value {
			subprefix := fmt.Sprintf("%s.%d", prefix, i)
			parseJSONValue(subvalue, subprefix, returnParams)
		}
	default:
		returnParams[prefix] = fmt.Sprintf("%v", value)
	}
}

func parseParams(p parser.HistoryItem) (string, string, EndpointParams) {

	if p.Params == "" {
		return "", "", EndpointParams{}
	}

	if p.Method == "GET" {
		e := EndpointParams{
			Params: Param{},
			Method: "GET",
		}

		e.Params = parseGETParams(p.Params)
		return p.Host, strings.Split(p.Path, "?")[0], e
	}

	if p.Method == "POST" {
		e := EndpointParams{
			Params: Param{},
			Method: "POST",
		}

		e.Params = parsePOSTParams(p.Params, p.ReqContentType)
		return p.Host, strings.Split(p.Path, "?")[0], e
	}

	return "", "", EndpointParams{}

}
