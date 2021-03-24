// byte, rune to represent characters

package main

import "fmt"

func main() {
	s1 := "I lke \n Go!" // raw string
	fmt.Println(s1)
	fmt.Println("He says: \"Hello!\"")
	fmt.Println(`He says: "Hello!"`)
}
