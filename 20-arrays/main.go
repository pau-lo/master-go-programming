// Array has fixed lenght, same type, indexable, like dicts in py

package main

import "fmt"

func main() {
	// array called accounts consisting of 3 integers

	accounts := [3]int{1, 2, 3}
	fmt.Println(accounts)

	friends := [4]string{"Dan", "Elma", "Paul", "Harriet"}
	fmt.Printf("%#v\n", friends)
	fmt.Println("length of accoutns: ", len(accounts))
	fmt.Println("length of friends: ", len(friends))

	a := [...]int{
		1,
		2,
		3,
		4,
		5, // this comma is mandatory
	}

	fmt.Println(a)

}
