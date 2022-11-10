package main

import "fmt"

func main() {
	fmt.Println(facultyRecursive(5))
}

func facultyRecursive(number int) int {
	if number < 1 {
		return 1
	}

	return number * facultyRecursive(number-1)
}
