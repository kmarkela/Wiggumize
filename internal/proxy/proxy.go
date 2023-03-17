package proxy

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func (p *Proxy) populateHistory(req *http.Request, res *http.Response) {
	log.Printf("Request: %+v\n", req)
	log.Printf("Response: %+v\n", res)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
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
		// p.handleHTTPSRequest(conn, req)
	} else {
		err := p.forwardHTTP(req, &conn)
		if err != nil {
			log.Println("Error forwarding HTTP request:", err)
			return
		}

	}
}

func (p *Proxy) forwardHTTP(req *http.Request, conn *net.Conn) error {
	var client *http.Client

	if p.UpstreamProxy != "" {
		proxyUrl, err := url.Parse(p.UpstreamProxy)
		if err != nil {
			return err
		}

		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	// TODO: error handling
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	p.populateHistory(req, res)
	res.Write(*conn)
	return nil
}

// func (p *Proxy) handleHTTPSRequest(conn net.Conn, req *http.Request) {
// 	host, _, err := net.SplitHostPort(req.Host)
// 	if err != nil {
// 		log.Println("Error splitting host and port:", err)
// 		return
// 	}

// 	// Generate a new self-signed certificate
// 	cert, err := p.generateSelfSignedCert(host)
// 	if err != nil {
// 		log.Println("Error generating certificate:", err)
// 		return
// 	}

// 	// Set up the TLS configuration for the server
// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{cert},
// 		ClientAuth:   tls.NoClientCert,
// 	}

// 	// Send a 200 OK response to the client to establish the SSL connection
// 	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

// 	// Create a new TLS connection with the client and the server
// 	tlsConn := tls.Server(conn, config)
// 	err = tlsConn.Handshake()
// 	if err != nil {
// 		log.Println("Error performing TLS handshake with client:", err)
// 		return
// 	}

// 	// Create a new request using the TLS connection
// 	req = &http.Request{
// 		Method: http.MethodConnect,
// 		Host:   req.Host,
// 	}

// 	// Connect to the server using the TLS connection
// 	serverConn, err := net.Dial("tcp", req.Host)
// 	if err != nil {
// 		log.Println("Error connecting to server:", err)
// 		return
// 	}
// 	defer serverConn.Close()

// 	// Set up the TLS configuration for the server
// 	serverConfig := &tls.Config{
// 		ServerName: host,
// 	}

// 	// Create a new TLS connection with the server
// 	serverTLSConn := tls.Client(serverConn, serverConfig)
// 	err = serverTLSConn.Handshake()
// 	if err != nil {
// 		log.Println("Error performing TLS handshake with server:", err)
// 		return
// 	}

// 	// Forward data between the client and the server
// 	go io.Copy(serverTLSConn, tlsConn)
// 	io.Copy(tlsConn, serverTLSConn)
// }

// func (p *Proxy) generateSelfSignedCert(host string) (tls.Certificate, error) {
// 	// Generate a new private key
// 	key, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		return tls.Certificate{}, err
// 	}

// 	// Create a self-signed certificate template
// 	template := x509.Certificate{
// 		SerialNumber:          big.NewInt(1),
// 		Subject:               pkix.Name{CommonName: host},
// 		NotBefore:             time.Now(),
// 		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
// 		BasicConstraintsValid: true,
// 	}

// 	// Create the certificate using the private key and the template
// 	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
// 	if err != nil {
// 		return tls.Certificate{}, err
// 	}

// 	// Encode the certificate and the private key in PEM format
// 	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
// 	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

// 	// Create the TLS certificate using the certificate and the private key
// 	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
// 	if err != nil {
// 		return tls.Certificate{}, err
// 	}

// 	return tlsCert, nil
// }
