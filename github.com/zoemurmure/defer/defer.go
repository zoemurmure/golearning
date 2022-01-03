package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//for i := 0; i < 4; i++ {
	//	defer fmt.Println("defer: ", -i)
	//	fmt.Println(i)
	//}
	newfile, error := os.Create("learngo.txt")
	if error != nil {
		fmt.Println("can't create file")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "test"); error != nil {
		fmt.Println("can't write file")
		return
	}

	newfile.Sync()
}
