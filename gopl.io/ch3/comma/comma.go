package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: comma.exe 19421789")
	}
	test := os.Args[1]
	fmt.Println(comma2(test))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + comma(s[n-3:])
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	// 如果是小数
	if idx := strings.LastIndex(s, "."); idx > 0 {
		n = idx
	}
	// 如果存在正负号
	if !unicode.IsDigit(rune(s[0])) {
		buf.WriteByte(s[0])
		s = s[1:]
		n -= 1
	}

	// 写入不足3位的部分
	i := n % 3
	buf.WriteString(s[:i])
	// 以三个数字为单位写入
	for i < n {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}
	// 写入小数部分
	buf.WriteString(s[n:])
	return buf.String()
}
