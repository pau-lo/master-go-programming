package main

import "fmt"

func main() {
	// you can provide type or not
	var a = 5   // int
	var b = 5.2 // float

	// a = b // cant assign float to int
	a = int(b)
	fmt.Println(a, b)

	// declare var of type int
	// var x int
	// 5 == a string literal (the value itself. doesn't have a name))
	// x = "5"

	// zero values
	// ensures values always holds a well defined value of its types
	// num types: 0
	// bool type: false
	// string type: "" (empty)
	// pointer type: nil

	var value int
	var prince float64
	var name string
	var done bool
	fmt.Println(value, prince, name, done)

}
