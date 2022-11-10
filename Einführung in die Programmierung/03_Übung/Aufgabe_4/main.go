package main

import "fmt"

func main() {
	fmt.Println(swapNumbers(4, 10))
}

func swapNumbers(num1, num2 int) (int, int) {
	return num2, num1
}
