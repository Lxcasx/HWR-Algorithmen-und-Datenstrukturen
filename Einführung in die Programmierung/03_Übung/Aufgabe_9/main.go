package main

import (
	"log"
	"time"
)

func main() {
	array := getArray()

	start := time.Now()

	for i := 0; i < 10000; i++ {
		doSomething(array)
	}

	elapsed := time.Since(start)
	log.Printf("faculty took %s", elapsed)
}

type Number struct {
	number int
	square int
}

func getArray() []Number {
	var array []Number = make([]Number, 100000)

	for i := 0; i < len(array); i++ {
		array[i] = Number{
			number: i,
			square: i * i,
		}
	}

	return array
}

func doSomething(array []Number) {
	test := array
	_ = test
	//fmt.Println(test)
}
