package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(num int) int {
	return num - 1
}

// Sum two integer number
func Sum(num1, num2 int) int {
	return num1 + num2
}
