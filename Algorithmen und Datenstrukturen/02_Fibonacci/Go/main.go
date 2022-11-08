package main

import "fmt"

func main() {
	fibonacciNumber := getFibonacciNumAt(7)
	fmt.Println(fibonacciNumber)
}

func getFibonacciNumAt(index int) int {
	nums := []int{0, 1}

	last := nums[0]
	last2 := nums[1]

	for i := 2; i < index+1; i++ {
		nums = append(nums[:], last+last2)
		last = last2
		last2 = nums[i]
	}

	return nums[index]
}
