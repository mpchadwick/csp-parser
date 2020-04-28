package main

import (
	"flag"
	"fmt"
	parser "github.com/mpchadwick/csp-parser/src"
	"os"
)

var url string

func init() {
	flag.StringVar(&url, "url", "", "The URL to check")
	flag.Parse()
}

func main() {
	fmt.Println(parser.FromUrl(url))
	os.Exit(0)
}