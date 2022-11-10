package main

import "fmt"

func main() {
	fmt.Println(getLargerNumber(15, 10))
}

func getLargerNumber(num1, num2 int) int {
	if num2 > num1 {
		return num2
	}

	return num1
}
