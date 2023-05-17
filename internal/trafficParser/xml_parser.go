package parser

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
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

func getParams(i Item) string {
	// divide request for headers and params

	var parts []string

	switch i.Method {
	case "POST":
		parts = strings.Split(i.Request.decodeBase64(), "\r\n\r\n")
	case "GET":
		parts = strings.Split(i.Path, "?")
	default:
		return ""
	}

	if len(parts) < 2 {
		return ""
	}
	// fmt.Println(parts[1])

	decodedString, err := url.QueryUnescape(parts[1])
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return ""
	}

	return decodedString

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
		history.RequestsList = append(history.RequestsList, HistoryItem{
			Time:     item.Time,
			URL:      item.URL,
			Host:     host,
			Path:     item.Path,
			Method:   item.Method,
			Request:  item.Request.decodeBase64(),
			Status:   item.Status,
			MimeType: item.MimeType,
			Response: item.Response.decodeBase64(),
			Params:   getParams(item),
		})
		history.ListOfHosts.Add(host)

	}

	return nil
}
