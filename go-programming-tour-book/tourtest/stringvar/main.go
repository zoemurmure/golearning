package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "go语言编程练习", "帮助信息")
	flag.StringVar(&name, "n", "go语言编程练习", "帮助信息")
	flag.Parse()

	fmt.Println(name)
}
