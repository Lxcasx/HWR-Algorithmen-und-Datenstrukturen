package main

import "fmt"

func main() {
	fmt.Println(faculty(5))
	fmt.Println(facultyRecursive(5))
}

func faculty(number int) int {
	num := 1

	for i := 1; i <= number; i++ {
		num = num * i
	}

	return num
}

func facultyRecursive(number int) int {
	if number < 1 {
		return 1
	}

	return number * facultyRecursive(number-1)
}
