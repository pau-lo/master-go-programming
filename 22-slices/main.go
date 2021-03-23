package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// A slice literal is declared just like an array literal,
	// except you leave out the element count:
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	numbers = append(numbers, 7)
	fmt.Println(numbers)
	fmt.Println(numbers[2:])

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"

	fmt.Println(cars, newCars)

	// checking size for slice and array
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	fmt.Println("Array size: ", unsafe.Sizeof(a))
	fmt.Println("Slice size: ", unsafe.Sizeof(b))

}
