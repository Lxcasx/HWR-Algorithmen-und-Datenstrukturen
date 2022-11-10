package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		for k := 0; k < 13; k++ {
			var num int

			if i == 0 {
				num = k
			} else if k == 0 {
				num = i
			} else {
				num = i * k
			}

			fmt.Printf("%03d ", num)
		}

		fmt.Println()
	}
}
