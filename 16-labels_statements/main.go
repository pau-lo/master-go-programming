package main

import "fmt"

func main() {
	outer := 19
	_ = outer
	people := [3]string {"John", "Erick", "Mel"}
	friends := [2]string {"Gio", "Paulo"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
 			}
		}
		fmt.Println("Next instruction after the break")
	}

}
