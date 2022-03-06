package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//s := `<p>Links:</p><ul><script>alert('123')</script><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	//doc, err := html.Parse(strings.NewReader(s))
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sameelement: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}
	//fmt.Println(n.Type)
	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if s != "" {
			fmt.Println(s)
		}
	}
	if !(n.Data == "script" || n.Data == "style" || n.Data == "noscript") {
		visit(n.FirstChild)
	}
	visit(n.NextSibling)

}
