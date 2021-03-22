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
}
