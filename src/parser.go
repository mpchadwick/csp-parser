package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FromUrl(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//fromHeaders := fromHeaders(resp)

	fromMeta := fromMeta(resp)

	return fromMeta
}

func fromHeaders(resp *http.Response) (string) {
	value := resp.Header.Get("Content-Security-Policy")
	
	if value == "" {
		value = resp.Header.Get("Content-Security-Policy-Report-Only")
	}

	return value
}

func fromMeta(resp *http.Response) (string) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	value := ""

	tags := doc.Find("meta")
	tags.Each(func(i int, s *goquery.Selection) {
		httpEquiv, _ := s.Attr("http-equiv")
		if httpEquiv == "Content-Security-Policy" {
			content, _ := s.Attr("content")
			value = content
		}
	})

	return value
}