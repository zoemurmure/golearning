package main

import (
	"os"
	"strings"

	"gopl.io/ch4/poster"
)

func main() {
	title := strings.Join(os.Args[1:], " ")
	if err := poster.Query(title); err != nil {
		panic(err)
	}
}
