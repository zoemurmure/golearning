package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("test1"))
	c2 := sha256.Sum256([]byte("test2"))
	fmt.Printf("%08b\n%08b\n%d\n", c1, c2, differentBit(c1, c2))
}

func differentBit(c1, c2 [32]byte) int {
	var result int
	for i := 0; i < 32; i++ {
		result += int(pc[c1[i]^c2[i]])
	}

	return result
}
