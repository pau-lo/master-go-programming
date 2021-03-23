package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("%v\n", numbers)

	//numbers[0] = 7

	for i, v := range numbers {
		fmt.Println("index:", i, "value: ", v)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index", i, "value:", numbers[i])
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, " value:", numbers[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	// keyed elements
	grades := [4]int{
		0: 10,
		1: 7,
		2: 8,
		3: 9,
	}
	fmt.Println("grades:", grades)

}
