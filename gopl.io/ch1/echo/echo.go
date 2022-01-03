package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//var name string
	//name = os.Args[0]
	//fmt.Println("Program:", name)
	//for idx, param := range os.Args[1:] {
	//	fmt.Println(idx, ":", param)
	//}

	var sep, params string
	start := time.Now()
	for i := 0; i < 10000; i++ {
		for _, param := range os.Args[1:] {
			params += sep + param
			sep = " "
		}
	}

	//fmt.Println(params)
	sec := time.Since(start).Milliseconds()
	fmt.Println("for:", sec)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		params = strings.Join(os.Args[1:], " ")
	}
	sec = time.Since(start).Milliseconds()
	fmt.Println("join:", sec)
}
