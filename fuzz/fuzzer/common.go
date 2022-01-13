package fuzzer

import (
	"log"
	"os"
)

var (
	iteration int = 10000
	cmd       []string
	method    int = -1
)
var Funcs = []fuzzFunc{byteOverwrite, magic}

func SetConfig(iter int, m int, c []string) {
	iteration = iter
	method = m
	cmd = c
}

// check if err is nil
func checkError(pref string, err error) {
	if err != nil {
		log.Fatal(pref + ": " + error.Error(err))
	}
}

// createa mutated.jpg with parameter
func createNew(data []byte) {
	file, err := os.OpenFile("mutated.jpg", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	checkError("Open mutate file", err)

	_, err = file.Write(data)
	checkError("Write mutate file", err)

	err = file.Close()
	checkError("Close mutate file", err)
}
