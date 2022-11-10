package main

import "fmt"

func main() {
	printLargerstNumber(5, 4)
}

func printLargerstNumber(num1, num2 int) {
	largest := num1

	if num2 > num1 {
		largest = num2
	}

	fmt.Println(largest)
}
