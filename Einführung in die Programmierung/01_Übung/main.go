package main

import (
	"fmt"
	"math"
)

func main() {
	/* Zusatz */
	res := math.Sqrt(3) * math.Sqrt(3)
	fmt.Println(res)

	res = 1.0 / 49 * 49
	fmt.Println(res)

	/* Zusatz */
	/*var i int64 = 2333234234234234423
	fmt.Println(unsafe.Sizeof(i))*/

	/* Aufgabe 5 */
	/*i, _ := strconv.Atoi("122")
	fmt.Println(i)

	s := strconv.Itoa(122)
	fmt.Println(s)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("15.4", 32)
	fmt.Println(f)

	ie, _ := strconv.ParseInt("123", 0, 16)
	fmt.Println(ie)

	fb := strconv.FormatBool(true)
	fmt.Println(fb)

	ff := strconv.FormatFloat(15.4, 32, 32, 32)
	fmt.Println(ff)*/

	/* Aufgabe 4 */
	/*var hello string = "Hallo Welt"
	var random string = "ÄÖÜaöü"

	fmt.Println(hello)
	fmt.Println(random)

	fmt.Println(len(hello))
	fmt.Println(len(random))

	fmt.Println(hello[0:5])
	fmt.Println(random[0:1])*/

	/* Aufgabe 3 */
	/*var num16 int16 = 30000
	var num32 int32 = 40000

	fmt.Println(int32(num16))
	fmt.Println(int16(num32))*/

	/* Aufgabe 2 */
	/*const pi = 3.14159265359

	var num float32 = 2 * pi
	var num2 float64 = 2 * pi

	fmt.Println(num)
	fmt.Println(num2)*/
}
