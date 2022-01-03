package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: printsha.exe [2/3/5] string")
	}
	if len(os.Args) == 2 {
		fmt.Printf("sha256 of %s is %x\n", os.Args[1], sha256.Sum256([]byte(os.Args[1])))
	}
	switch os.Args[1] {
	case "2":
		fmt.Printf("sha256 of %s is %x\n", os.Args[2], sha256.Sum256([]byte(os.Args[2])))
	case "3":
		fmt.Printf("sha384 of %s is %x\n", os.Args[2], sha512.Sum384([]byte(os.Args[2])))
	case "5":
		fmt.Printf("sha512 of %s is %x\n", os.Args[2], sha512.Sum512([]byte(os.Args[2])))
	default:
		panic("usage: printsha.exe [2/3/5] string")
	}
}
