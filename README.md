# Wiggumize

> Burp Suite History Analysis Tool

Wiggumize is a tool developed in Golang for analyzing Burp Suite history files in XML format. It aims to assist security researchers in identifying potential security issues and providing tips on where to focus their investigations. The tool performs various checks on the Burp Suite history data and extracts valuable insights to aid in the security assessment process.

## Features

Wiggumize currently supports the following checks:

- **SSRF Detection**: This module identifies URLs present in request parameters, helping to uncover potential Server-Side Request Forgery (SSRF) vulnerabilities.

- **404 Detection**: The tool searches for 404 messages in responses, indicating possible misconfigurations or vulnerabilities in the web application's hosting.

- **XML Analysis**: This module scans request parameters for XML data, allowing security researchers to identify XML-related vulnerabilities.

- **Redirect Detection**: Wiggumize is capable of identifying redirects, which could lead to open redirection vulnerabilities.

- **Secrets Detection**: The secrets detection module focuses on identifying sensitive information like API keys within the Burp Suite history.

- **LFI Indication**: By analyzing filenames in request parameters, the tool gives indications of possible Local File Inclusion (LFI) vulnerabilities.

- **Parameter Parsing**: Wiggumize parses GET and POST (JSON) parameters for further analysis, enhancing the overall security assessment process.

## Install 

1. `go get`
```bash
go get github.com/kmarkela/Wiggumize
```

2. build from sorces
```bash
git clone https://github.com/kmarkela/Wiggumize.git
cd Wiggumize
go build
```

## Usage

The tool provides a command-line interface with the following parameters:

```shell
-f   Path to XML file with Burp history
-o   Path to output file (default: report.md)
-a   Action. 'scan' for history analysis (default), 'search' for pattern search
```

### Scan Mode (Default)

In scan mode, Wiggumize analyzes the provided Burp history XML file and performs various security checks. The results are compiled into a Markdown report (by default, `report.md`).

### Search Mode

Search mode enables researchers to search for specific patterns within the Burp history. Regular expressions can be used to search across different fields, including request method, headers, content type, body, response headers, response content type, and response body.

#### Search Example:

Search for requests with the following criteria:
- Request method is POST
- Request body contains the string "admin"
- Response content type is not HTML
- Response body contains the string "success"

```shell
ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success
```

## Contributing

Contributions to Wiggumize are welcome! If you have suggestions for new checks, improvements, or bug fixes, please open an issue or submit a pull request.

---

Wiggumize empowers security researchers to uncover potential vulnerabilities and enhance their web application assessments using Burp Suite history. Its modular design and extensibility make it a valuable tool in the arsenal of any security professional. Feel free to explore, contribute, and make the web a safer place with Wiggumize!