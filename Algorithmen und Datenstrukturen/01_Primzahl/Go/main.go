package main

import "fmt"

func main() {
	number := 5
	isPrime := isNumPrime(number)

	fmt.Printf("Is prime number: %t", isPrime)
}

func isNumPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
