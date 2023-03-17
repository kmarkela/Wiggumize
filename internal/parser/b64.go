package parser

import "encoding/base64"

func B64Decode(s string) string {
	// TODO: handle errors
	ns, _ := base64.StdEncoding.DecodeString(s)
	return string(ns)
}
