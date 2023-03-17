package proxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
)

type ProxyHistory struct {
	Hosts []string
	Items []Item
}
type Item struct {
	Time     string
	URL      string
	Host     string
	Port     string
	Protocol string
	Method   string
	Path     string
	Request  string
	Status   string
	MimeType string
	Response string
}

type Proxy struct {
	// Upstream proxy address
	UpstreamProxy string

	// Certificate and key for MITM decryption
	CertFile string
	KeyFile  string
	History  *ProxyHistory
}

type RequestResponse struct {
	Request  *http.Request
	Response *http.Response
}

func (p *Proxy) populateHistory(r *RequestResponse) {
	log.Printf("Request: %+v\n", r.Request)
	log.Printf("Response: %+v\n", r.Response)

	// Read the response body
	body, err := io.ReadAll(r.Response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s\n", err)
	} else {
		log.Printf("Response body: %s\n", string(body))
	}

	// Add the request/response to the history
	// if p.History != nil {
	// 	p.History.AddRequestResponse(r)
	// }
}

func (p *Proxy) Start(port int) {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

	log.Printf("Proxy server listening on port %d", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}

		go p.handleConnection(conn)
	}
}
func (p *Proxy) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Wrap the connection in a bufio reader for reading the request
	reader := bufio.NewReader(conn)

	// Parse the incoming request using http.ReadRequest
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	// We can't have this set.
	req.RequestURI = ""

	// Check if the request is HTTPS
	if req.Method == http.MethodConnect {
		log.Printf("HTTPS")
	} else {
		res, err := p.forwardHTTP(req, &conn)
		if err != nil {
			log.Println("Error forwarding HTTP request:", err)
			return
		}
		p.populateHistory(&RequestResponse{req, res})
	}
}

func (p *Proxy) forwardHTTP(req *http.Request, conn *net.Conn) (*http.Response, error) {
	var client *http.Client

	if p.UpstreamProxy != "" {
		proxyUrl, err := url.Parse(p.UpstreamProxy)
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	// TODO: error handling
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()

	// Write the response headers to the client
	writer := bufio.NewWriter(*conn)
	if err := res.Write(writer); err != nil {
		log.Println("Error writing response headers:", err)
		return res, err
	}

	// Write the response body to the client
	if _, err := io.Copy(writer, res.Body); err != nil {
		log.Println("Error writing response body:", err)
		return res, err
	}

	// // Close the response body after all data has been copied to the client
	// if err := res.Body.Close(); err != nil {
	// 	log.Println("Error closing response body:", err)
	// 	return res, err
	// }

	// Flush the writer to ensure all data is sent to the client
	if err := writer.Flush(); err != nil {
		log.Println("Error flushing writer:", err)
		return res, err
	}

	return res, nil
}
