package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go编程练习", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "go编程练习", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("%s", name)
}
