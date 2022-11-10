package main

import (
	"log"
	"time"
)

const amount int = 10000

func main() {
	/* fibonacci */
	start := time.Now()

	for i := 0; i < amount; i++ {
		_ = calcFibonacci(5)
	}

	elapsed := time.Since(start)
	log.Printf("faculty took %s", elapsed)

	/*  fibonacci recursive */
	start = time.Now()

	for i := 0; i < amount; i++ {
		_ = calcFibonacciRecursive(5)
	}

	elapsed = time.Since(start)
	log.Printf("facultyRecursive took %s", elapsed)
}

func calcFibonacci(index int) int {
	var numbers []int
	numbers = append(numbers, 0, 1)

	last := numbers[0]
	last2 := numbers[1]

	for i := 1; i < index+1; i++ {
		numbers = append(numbers, last+last2)

		last = last2
		last2 = numbers[i+1]
	}

	return numbers[index]
}

func calcFibonacciRecursive(index int) int {
	if index == 1 || index == 2 {
		return 1
	}

	return calcFibonacciRecursive(index-1) + calcFibonacciRecursive(index-2)
}
