package main

import "fmt"

func main() {
	// set i to 0, less than 10, increment
	// intialization, boolean condition, post statement
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// no while loop
	// but this is similar to it
	// j := 10 // initialization
	// for j >= 0 {
	// 	fmt.Println(j)
	// 	j-- // post statement
	// }

	// infinite loop
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum) // never reached

	for i := 0; i < 10; i++ {
		if i%2 !=0 {
			continue
		}
		fmt.Println(i) // if 'i' is an odd condition will not be met
	}
}
