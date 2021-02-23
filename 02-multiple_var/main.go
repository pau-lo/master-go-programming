package main

import "fmt"

func main() {

	// declaring two var
	car, cost := "Audi", 55000
	fmt.Println(car, cost)

	// declaring one new var with an existing one
	car, year := "BMW", 2018
	// muting the new car var error
	_ = year

	// normal declaration
	var opened = false
	// short declaration
	opened, file := true, "a.txt"

	_, _ = opened, file

	// multiple declaration
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// other way conscise way
	var a, b, c int
	fmt.Println(a, b, c)

	// Note:
	// short declaration := --> when we know the initial value
	// normal declaration = --> for multiple declaration

	// multiple assignments
	// declar i, j
	var i, j int

	// assign i, j
	i, j = 5, 8

	// again by swapping - reassigning
	j, i = i, j

	// muting erros
	_, _ = i, j

	fmt.Println(i, j)

	// expression w/ short declaration
	sum := 5 + 3
	fmt.Println(sum)

}
