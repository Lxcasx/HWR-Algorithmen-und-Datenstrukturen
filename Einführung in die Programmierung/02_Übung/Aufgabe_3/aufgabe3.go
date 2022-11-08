package main

import "fmt"

// Geben Sie das Array mit den 4 Schleifenarten aus.
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, e := range arr {
		fmt.Println(e)
	}
}
