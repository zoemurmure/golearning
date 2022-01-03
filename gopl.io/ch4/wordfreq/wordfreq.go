package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]float64)
	sum := 0.0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
		sum++
	}
	fmt.Println(sum)
	for k, v := range count {
		//fmt.Printf("%s\t%.2f\n", k, v)
		count[k] = v / sum
		fmt.Printf("%s\t%.5f\n", k, count[k])
	}

	//fmt.Println(count)
}
