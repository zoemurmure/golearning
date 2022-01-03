package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: basename.exe teststring")
	}
	test := os.Args[1]
	fmt.Println(basename1(test))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	idx := strings.LastIndex(s, "/")
	s = s[idx+1:]
	if idx = strings.LastIndex(s, "."); idx > 0 {
		s = s[:idx]
	}
	return s
}
