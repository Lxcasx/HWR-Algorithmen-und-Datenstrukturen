package main

import (
	"log"
	"time"
)

func main() {
	/* facultiy */
	start := time.Now()

	for i := 0; i < 10000; i++ {
		_ = faculty(5)
	}

	elapsed := time.Since(start)
	log.Printf("faculty took %s", elapsed)

	/* facultiy recursive */
	start = time.Now()

	for i := 0; i < 10000; i++ {
		_ = facultyRecursive(5)
	}

	elapsed = time.Since(start)
	log.Printf("facultyRecursive took %s", elapsed)
}

func facultyRecursive(number int) int {
	if number < 1 {
		return 1
	}

	return number * facultyRecursive(number-1)
}

func faculty(number int) int {
	num := 1

	for i := 1; i <= number; i++ {
		num = num * i
	}

	return num
}
