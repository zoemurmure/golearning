package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s = [...]int{1, 2, 3, 4, 5}
	//reverse(&s)
	//fmt.Println(s)

	rotate(s[:], -1)
	fmt.Println(s)

	var string = []string{"123", "123", "123", "123", "346", "234", "234", "342"}
	string = removeRepeat(string)
	fmt.Println(string)

	var s2 = []byte("测试   这是一句  测试      aewf   ")
	fmt.Printf("%s\n", s2)
	s2 = removeRepeatSpace((s2))
	fmt.Printf("%s\n", s2)
	s2 = reverse(s2)
	fmt.Printf("%s\n", s2)
}

//func reverse(s *[5]int) {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//}

func reverse(bs []byte) []byte {
	runes := []rune(string(bs))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}

func rotate(s []int, n int) {
	for n < 0 {
		n += len(s)
	}
	j, temp := 0, s[0]
	for {
		temp, s[(j+n)%len(s)] = s[(j+n)%len(s)], temp
		if j = (j + n) % len(s); j == 0 {
			temp, s[(j+n)%len(s)] = s[(j+n)%len(s)], temp
			break
		}
	}
}

func removeRepeat(s []string) []string {
	var current = 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[current] {
			//copy(s[current+1:], s[i:])
			s[current+1] = s[i]
			current++
		}
	}
	return s[:current+1]
}

func removeRepeatSpace(s []byte) []byte {
	current := 1
	for i := 1; i < len(s); {
		s[current] = s[i]
		if unicode.IsSpace(rune(s[i])) {
			for i < len(s) && unicode.IsSpace(rune(s[i])) {
				i++
			}
			i--
		}
		current++
		i++
	}
	return s[:current]
}
