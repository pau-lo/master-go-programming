/*

if, else if and else statements that are used for decision making

syntax:

if some_condition_is_true {
	// executed code
} else if some_other_condition_is_true {
	//execute code
} else {
	// execute this code
}

*/

// Examples

package main

import "fmt"

func main() {

	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive!")
	}

	if price <= 100 && inStock == true {
		fmt.Println("Buy it!!!")
	}

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge!")
	} else {
		fmt.Println("It's Expensive")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("Sorry, you can't vote.  Please return in %d years !\n", age-18)
	} else if age == 18 {
		fmt.Printf("This is your first vote!")

	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's important.")
	} else {
		fmt.Printf("Invalid age!")
	}

}
