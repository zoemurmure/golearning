package main

import (
	"fmt"
	"math"
)

func main() {
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)
}
