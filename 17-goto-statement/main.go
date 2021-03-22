package main

import "fmt"

func main() {
	i := 0
loop:
	if i < 5 {
		i++
		goto loop // doesn't have any restriction
	}
	fmt.Printf("Value of i:", i)

	// Not allowed
	// 	goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("WHERE IS THIS GOING?")
}
