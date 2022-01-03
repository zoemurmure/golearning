package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	count := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		switch {
		case r == unicode.ReplacementChar:
			count["invalid"]++
		case unicode.IsNumber(r):
			count["number"]++
		case unicode.IsLetter(r):
			count["letter"]++
		case unicode.IsPunct(r):
			count["punct"]++
		case unicode.IsSpace(r):
			count["space"]++
		default:
			count["others"]++
		}
	}

	fmt.Println(count)
}
