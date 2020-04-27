package main

import (
	"bufio"
	"fmt"
	parser "github.com/mpchadwick/csp-parser/src"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		result := parser.Parse(text)

		fmt.Print(result)
		os.Exit(0)
	}
}