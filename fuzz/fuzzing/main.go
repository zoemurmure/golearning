package main

import (
	"fmt"
	"fuzz/fuzzer"
	"os"
	"time"
)

//var cmds = []string{"rm", "pr", "fi", "fc", "ex"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fuzzer.exe <valid_jpg>")
		os.Exit(1)
	}

	start := time.Now()

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("main: " + err.Error())
	}

	fuzzer.SetConfig(100000, 0, []string{"./vulfunc", "./mutated.jpg", "3"})
	//fuzzer.Debug("./vulfunc", []string{"./Canon_40D.jpg", "3"}, data)
	fuzzer.Run(data)

	duration := time.Since(start).Microseconds()
	fmt.Println("Execution time:", duration, "ms")
}
