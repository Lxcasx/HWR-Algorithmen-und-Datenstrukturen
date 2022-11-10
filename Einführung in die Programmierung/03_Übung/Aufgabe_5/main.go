package main

import "fmt"

func main() {
	fmt.Println(faculty(5))
}

func faculty(number int) int {
	num := 1

	for i := 1; i <= number; i++ {
		num = num * i
	}

	return num
}
