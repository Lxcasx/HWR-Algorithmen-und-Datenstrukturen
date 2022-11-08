package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := getRandomNumbers()

	fmt.Println(array)

	for i := 1; i < len(array); i++ {
		num1 := 0
		num2 := 0

		for k := i; k > 0; k-- {
			num1 = array[k-1]
			num2 = array[k]

			if num1 > num2 {
				array[k-1] = num2
				array[k] = num1
			}
		}
	}

	fmt.Println(array)
}

func getRandomNumbers() [50]int {
	array := [50]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(50-0) + 0
	}

	return array
}
