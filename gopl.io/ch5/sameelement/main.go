package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var Elements = map[string]int{}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sameelement: %v\n", err)
		os.Exit(1)
	}
	count(doc)
	fmt.Println(Elements)
}

func count(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		Elements[n.Data]++
	}

	count(n.FirstChild)
	count(n.NextSibling)
}
