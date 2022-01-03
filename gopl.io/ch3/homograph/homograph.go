package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 3 {
		panic("usage: homograph.exe test estt")
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Printf("%t\n", homograph([]byte(s1), []byte(s2)))
}

func homograph(s1, s2 []byte) bool {
	sort.SliceStable(s1, func(i, j int) bool { return s1[i] <= s1[j] })
	sort.SliceStable(s2, func(i, j int) bool { return s2[i] <= s2[j] })

	return bytes.Equal(s1, s2)
}
