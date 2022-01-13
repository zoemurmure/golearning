package fuzzer

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var baseData []byte
var succeedCount int

func Run(data []byte) {
	baseData = data
	for i := 0; i < iteration; i++ {
		mutateData := getData(baseData, method)
		createNew(mutateData)
		runCommand(i, mutateData)
		if i%1000 == 0 {
			fmt.Printf("%d/%d\r", i, iteration)
		}
	}
}

func runCommand(counter int, data []byte) {
	defer func() {
		handler := recover()
		if handler != nil {
			result := fmt.Sprintf("%v", handler)
			if strings.Contains(result, "123") {
				filename := fmt.Sprintf("crashes/crash_%d.jpg", counter)
				err := os.WriteFile(filename, data, 0644)
				checkError("Write crash file", err)
				return
			}
		}
	}()

	//cmd := exec.Command("./exif", "./mutated.jpg", "-verbose")
	//cmd := exec.Command("exiv2", ccmd, "-v", "./mutated.jpg")
	ccmd := exec.Command(cmd[0], cmd[1:]...)

	if output, err := ccmd.CombinedOutput(); err != nil {
		if strings.Contains(string(output), strconv.Itoa(succeedCount+1)+" succeed") {
			fmt.Println(string(output))
			baseData = data
			succeedCount++
		}
		panic("Run ccmd" + ": " + err.Error())
	}
}

func getData(data []byte, method int) []byte {
	var mutateData []byte
	switch method {
	case -1:
		start := time.Now()
		rand.Seed(start.UnixNano())
		mutateData = Funcs[rand.Intn(len(Funcs))](data)
	case 0:
		mutateData = Funcs[0](data)
	case 1:
		mutateData = Funcs[1](data)
	}
	return mutateData
}
