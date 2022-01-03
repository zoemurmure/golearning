package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch4/xkcd"
)

func main() {
	var num int
	var query string

	if len(os.Args) == 2 {
		query = os.Args[1]
	} else if len(os.Args) > 2 {
		num, _ = strconv.Atoi(os.Args[1])
		//if err != nil {
		//	panic("Usage: xkcdQuery.exe [#data] query_string")
		//}
		query = strings.Join(os.Args[2:], " ")
	} else {
		panic("Usage: xkcdQuery.exe [#data] query_string")
	}
	if num > 0 {
		xkcd.EstablishDatabase(num)
	}
	result, err := xkcd.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
