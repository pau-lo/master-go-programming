// scope --> visibility
/*
A name cannot be declared again in the same scope.
	3 types:
		1. File Scope
		2. Package Scope
		3. Block (local) Scope
*/

package main

import f "fmt" // alias are permitted

const done = false // package scope

func main() {
	x := 1 // block (local)  scope
	f.Println(x, done)
	f.Println(done)
	f1()
}

func f1() {
	const done = true
	f.Printf("Done is f1(): %v\n", done) // package scope
	// a int;
	a := 10
	_ = a
}
