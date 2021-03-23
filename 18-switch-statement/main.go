package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("Python")

	case "Go", "golang":
		fmt.Println("Good, golang.")

	default:

		fmt.Println("Any other programming is a good start!")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 == 0:
		fmt.Println("Even number!")
	default:
		fmt.Println("Never here!")
	}
}
