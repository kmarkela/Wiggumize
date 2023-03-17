package parser

import (
	"Wiggumize/utils"
	"encoding/xml"
	"io/ioutil"
	"net"
	"net/url"
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
	Request        Request  `xml: "request"`
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

type Response struct {
	Base64 bool   `xml:"base64,attr"`
	Value  string `xml:",chardata"`
}

// Parse is a method that parses an XML file.
func (p *XMLParser) Parse(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Use the xml.Unmarshal() function to parse the XML data.
	err = xml.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	return nil
}

// return list of uniq
func (p *XMLParser) ListOfHosts() []string {
	set := &utils.Set{}
	for _, item := range p.ItemElements {
		s := item.Protocol + "://" + item.Host.Value + ":" + item.Port
		set.Add(s)

	}

	return set.Keys()
}

type hostSt struct {
	protocol string
	host     string
	port     string
}

func (p *XMLParser) FilterByHost(URLs []string) {
	filteredItems := []Item{}

	parsedHosts := []hostSt{}
	for _, URL := range URLs {
		u, _ := url.Parse(URL)
		host, port, _ := net.SplitHostPort(u.Host)
		parsedHosts = append(parsedHosts, hostSt{u.Scheme, host, port})
	}

	for _, item := range p.ItemElements {
		for _, host := range parsedHosts {
			if item.Host.Value == host.host && item.Port == host.port && item.Protocol == host.protocol {
				filteredItems = append(filteredItems, item)
				break
			}
		}
	}
	p.ItemElements = filteredItems
}
