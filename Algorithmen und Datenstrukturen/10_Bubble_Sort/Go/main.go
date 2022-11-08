package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := getRandomNumbers()

	fmt.Println(array)

	for i := len(array); i > 1; i-- {
		for k := 0; k < i-1; k++ {
			num1 := array[k]
			num2 := array[k+1]

			if num1 > num2 {
				array[k] = num2
				array[k+1] = num1
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
