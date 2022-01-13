package fuzzer

import "math/rand"

type fuzzFunc func([]byte) []byte

const flipPercent = 0.02

var nflips int

// randomly change flipPercent of bytes in data
// randomly flip one bit in every byte
func bitFlip(data []byte) []byte {
	nflips = int((float64(len(data)) - 4) * flipPercent)
	result := make([]byte, len(data))
	if nbytes := copy(result, data); nbytes == 0 {
		panic("bitFlip: Error when copying")
	}

	mask := []byte{1, 2, 4, 8, 16, 32, 64, 128}
	var chosenIdx []int

	for i := 0; i < nflips; i++ {
		chosenIdx = append(chosenIdx, rand.Intn(len(data)-4)+2)
	}

	for _, x := range chosenIdx {
		flipIdx := rand.Intn(8)
		result[x] ^= mask[flipIdx]
	}

	return result
}

// randomly overwrite flipPercent of bytes in data
// randomly choose from [0, 256)
func byteOverwrite(data []byte) []byte {
	nflips = int((float64(len(data)) - 4) * flipPercent)
	result := make([]byte, len(data))
	if nbytes := copy(result, data); nbytes == 0 {
		panic("byteOverwrite: Error when copying")
	}

	var chosenIdx []int
	for i := 0; i < nflips; i++ {
		chosenIdx = append(chosenIdx, rand.Intn(len(data)-4)+2)
	}

	for _, x := range chosenIdx {
		result[x] = byte(rand.Intn(256))
	}

	return result
}

// randomly overwrite 1~4 bytes in data with magic number
func magic(data []byte) []byte {
	magicNum := [][]byte{
		{0xFF}, {0x7F}, {0x00},
		{0xFF, 0xFF}, {0x00, 0x00},
		{0xFF, 0xFF, 0xFF, 0xFF},
		{0x00, 0x00, 0x00, 0x00},
		{0x80, 0x00, 0x00, 0x00},
		{0x40, 0x00, 0x00, 0x00},
		{0x7F, 0xFF, 0xFF, 0xFF},
	}
	result := make([]byte, len(data))
	if nbytes := copy(result, data); nbytes == 0 {
		panic("magic: Error when copying")
	}

	chosenIdx := rand.Intn(len(data)-8) + 2
	chosenMagic := rand.Intn(len(magicNum))

	for _, i := range magicNum[chosenMagic] {
		result[chosenIdx] = byte(i)
		chosenIdx++
	}

	return result
}
