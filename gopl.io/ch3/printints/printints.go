package main

import (
	"bytes"
	"fmt"
)

func main() {
	test := []int{4, 2, 5, 6, 242, 2}
	fmt.Println(intsToString(test))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
