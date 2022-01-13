package main

import (
	"fmt"
	"os"
	"strconv"
)

var checkPos = []float64{0.33, 0.5, 0.67}
var checkVal = []byte{0x6c, 0x57, 0x21}

func main() {
	if len(os.Args) < 3 {
		panic("usage vulfunc.exe filename ncheck ([0,3])")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("main: error when reading file")
	}

	n := len(data)
	nCheck, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("main: error when convert string to number")
	}

	for i := 0; i < nCheck; i++ {
		if checkFunc(data, int(checkPos[i]*float64(n)), checkVal[i]) {
			fmt.Printf("[√]Check %d succeed!\n", i+1)
		} else {
			fmt.Printf("[x]Check %d failed!\n", i+1)
			os.Exit(1)
		}
	}
	vuln(data, n)
}

func checkFunc(data []byte, pos int, val byte) bool {
	//fmt.Printf("%x\n", pos)
	return data[pos] == val
}

func vuln(data []byte, n int) {
	defer func() {
		if handler := recover(); handler != nil {
			fmt.Println("ERROR: index out of range [20] with length 20")
			os.Exit(123)
		}
	}()

	fmt.Println("[√]Pass all checks!")
	newBuf := make([]byte, 20)

	for i := 0; i < n; i++ {
		newBuf[i] = data[i]
	}
}
