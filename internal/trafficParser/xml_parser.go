package parser

import (
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"regexp"
	"strings"
)

// XMLParser is a struct that represents the XML parser.
//type XMLParser struct{}

type XMLParser struct {
	XMLName      xml.Name `xml:"items"`
	BurpVersion  string   `xml:"burpVersion,attr"`
	ExportTime   string   `xml:"exportTime,attr"`
	ItemElements []Item   `xml:"item"`
}

type Item struct {
	Time           string   `xml:"time"`
	URL            string   `xml:"url"`
	Host           Host     `xml:"host"`
	Port           string   `xml:"port"`
	Protocol       string   `xml:"protocol"`
	Method         string   `xml:"method"`
	Path           string   `xml:"path"`
	Extension      string   `xml:"extension"`
	Request        Request  `xml:"request"`
	Status         string   `xml:"status"`
	ResponseLength string   `xml:"responselength"`
	MimeType       string   `xml:"mimetype"`
	Response       Response `xml:"response"`
	Comment        string   `xml:"comment"`
}

type Host struct {
	Value string `xml:",chardata"`
	IP    string `xml:"ip,attr"`
}

type Request struct {
	Base64 bool   `xml:"base64,attr"`
	Value  string `xml:",chardata"`
}

func (r Request) decodeBase64() string {

	// Return if not encoded
	if !r.Base64 {
		return r.Value
	}

	stringBytes, _ := base64.StdEncoding.DecodeString(r.Value)
	return string(stringBytes)
}

type Response struct {
	Base64 bool   `xml:"base64,attr"`
	Value  string `xml:",chardata"`
}

func (r Response) decodeBase64() string {

	// Return if not encoded
	if !r.Base64 {
		return r.Value
	}

	stringBytes, _ := base64.StdEncoding.DecodeString(r.Value)
	return string(stringBytes)
}

// Parse is a method that parses an XML file.
func (p *XMLParser) Parse(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// TODO: takes a few seconds! needs to be looked at
	// Use the xml.Unmarshal() function to parse the XML data.
	err = xml.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	return nil
}

// TODO: move to utils

type reqRes struct {
	reqBody    string
	reqHeaders string
	resBody    string
	resHeaders string
	reqParams  string
}

func parseReqRes(i Item) reqRes {

	var reqParts []string
	var resParts []string

	switch i.Method {
	case "POST", "PUT", "PATCH":
		reqParts = strings.Split(i.Request.decodeBase64(), "\r\n\r\n")
	case "GET":
		reqParts = strings.Split(i.Path, "?")
	}

	resParts = strings.Split(i.Response.decodeBase64(), "\r\n\r\n")

	var rr reqRes = reqRes{
		resHeaders: resParts[0],
		reqHeaders: reqParts[0],
	}

	if len(resParts) > 1 {
		rr.resBody = resParts[1]
	}

	if len(reqParts) > 1 && i.Method == "GET" {
		//todo: refactor this. Parametes is used in Scaner.
		rr.reqParams = reqParts[1]
	} else if len(reqParts) > 1 {
		rr.reqBody = reqParts[1]
	}

	return rr
}

// func getParams(i Item) string {
// 	// divide request for headers and params

// 	var parts []string

// 	switch i.Method {
// 	case "POST":
// 		parts = strings.Split(i.Request.decodeBase64(), "\r\n\r\n")
// 	case "GET":
// 		parts = strings.Split(i.Path, "?")
// 	default:
// 		return ""
// 	}

// 	if len(parts) < 2 {
// 		return ""
// 	}
// 	// fmt.Println(parts[1])

// 	// Get elements from index 1 till the end of the list. (remove headers)
// 	elements := parts[1:]

// 	// Create a string by joining the elements with a separator
// 	result := strings.Join(elements, " ")

// 	if i.Method == "POST" {
// 		return result
// 	}

// 	decodedString, err := url.QueryUnescape(result)
// 	if err != nil {
// 		fmt.Println("Error decoding URL:", err)
// 		return ""
// 	}

// 	return decodedString

// }

// TODO: move to utils
func getContentType(headerString string) string {
	lines := strings.Split(headerString, "\n")
	contentTypeRegex := regexp.MustCompile(`Content-Type:\s*(.*)`)

	for _, line := range lines {
		match := contentTypeRegex.FindStringSubmatch(line)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		}
	}

	return ""
}

func (p *XMLParser) PopulateHistory(file string, history *BrowseHistory) error {
	// Parsing XML history and populating BrowseHistory struct

	// Parser XML file and polulate XMLParser
	err := p.Parse(file)
	if err != nil {
		return err
	}

	// Populate BrowseHistory
	for _, item := range p.ItemElements {
		host := item.Protocol + "://" + item.Host.Value + ":" + item.Port

		ReqContentType := ""
		if item.Method == "POST" {
			ReqContentType = getContentType(item.Request.decodeBase64())
		}

		var rr reqRes = parseReqRes(item)

		history.RequestsList = append(history.RequestsList, HistoryItem{
			Time:           item.Time,
			URL:            item.URL,
			Host:           host,
			Path:           item.Path,
			Method:         item.Method,
			Request:        item.Request.decodeBase64(),
			Status:         item.Status,
			ReqContentType: ReqContentType,
			ResContentType: getContentType(item.Response.decodeBase64()),
			Response:       item.Response.decodeBase64(),
			ReqHeaders:     rr.reqHeaders,
			ResHeaders:     rr.resHeaders,
			ReqBody:        rr.reqBody,
			ResBody:        rr.resBody,
			Params:         rr.reqParams,
		})

		history.ListOfHosts.Add(host)

	}

	return nil
}
