/*

data types:

	bool ==> true or false

	string ==> char enclosed in quotes

	array ==> number seq. of elements of a single type, fixed length.

	slice ==> dynamic, can grow or shrink

	map ==> unordered group of elements of on type, indexed by a set of
			uniques keys of another type --> similar to a dict in python

	struct type ==> seq. of named elements, called fields, each of which has a name and type.
					can be compare to class in oop.

	pointer type ==> a var. that stores the memory address of another variable
					 The value of an uninitialized pointer is nil.

	Function and interface type ==>

	Channel type ==> provides a mechanism for concurrently executing functions to communicate by
					 sending and receiving values of a specified element type.
*/

package main

import "fmt"

func main() {
	var b bool = true
	// bool type
	fmt.Printf("%T\n", b)

	// string type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	// array type
	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", number)

	// slice type
	var cities = []string{"San Francisco", "Los Angeles", "New York", "Seattle", "Boston"}
	fmt.Printf("%T\n", cities)

	// map type
	balances := map[string]float64{
		"usd": 100.5,
		"eur": 25.5,
	}
	fmt.Printf("%T\n", balances)

	// struct type
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// pointer type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n: ", ptr, ptr)

	// function type
	fmt.Printf("%T\n", f)
}

// function f
func f() {

}
