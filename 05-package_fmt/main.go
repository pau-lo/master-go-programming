package main

import "fmt"

func main() {
	name := "Paulin"
	fmt.Println("Hey, I am", name+".")

	a, b := 4.0, 7.0

	fmt.Println("Sum", a+b, "Mean Value", (a+b)/2)

	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.1459

	fmt.Printf("Figure is a %s\n", figure)
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n",
		figure, radius, float64(radius)*2*pi)

	// %q for quoted string
	fmt.Printf("This is %q\n", figure)

	// /&v -> can be replace with any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n",
		figure, radius, float64(radius)*2*pi)

	// %T -> prints out the type of var
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	// %t -> bool
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b -> base 2
	fmt.Printf("%b \n", 55)

}
