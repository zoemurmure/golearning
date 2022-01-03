package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

type popcountFunc func(uint64) int

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCountBit(x uint64) int {
	var result int
	for x != 0 {
		if x&1 == 1 {
			result++
		}
		x >>= 1
	}
	return result
}

func PopCountClear(x uint64) int {
	var result int
	for x&(x-1) != x {
		result++
		x = x & (x - 1)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		panic("usage: popcound.exe [num]")
	}
	num, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%064b\n", num)

	fmt.Println(escape(PopCount, num))

	fmt.Println(escape(PopCountBit, num))
	fmt.Println(escape(PopCountClear, num))
	fmt.Println(escape(PopCountLoop, num))
}

func escape(f popcountFunc, num uint64) int64 {
	start := time.Now()

	for i := 0; i < 100000; i++ {
		f(num)
	}

	t := time.Since(start).Microseconds()
	return t
}
