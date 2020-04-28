package parser

import (
	"log"
	"net/http"
)

func FromUrl(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fromHeaders := fromHeaders(resp)

	return fromHeaders
}

func fromHeaders(resp *http.Response) (string) {
	value := resp.Header.Get("Content-Security-Policy")
	
	if value == "" {
		value = resp.Header.Get("Content-Security-Policy-Report-Only")
	}

	return value
}