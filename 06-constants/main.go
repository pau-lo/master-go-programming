package main

import "fmt"

func main() {

	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration ins seconds: %v\n", duration*secondsInHour)

	// This will give an run time  error
	x, y := 5, 1 //0
	fmt.Println(x / y)

	// const belongs to compile time so errors are discovered earlier
	// const a = 5
	// const b = 0

	// fmt.Println(a / b)

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	// num2, num3 will get their type from num1
	// because in a group constant, the values can repeat
	const (
		num1 = 250
		num2
		num3
	)

	fmt.Println(num1, num2, num3)
}
