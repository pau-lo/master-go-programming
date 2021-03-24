package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	word := strings.Contains("I am enjoying Golang!", "Gol")
	p(word)

	word = strings.ContainsAny("success", "xyz")
	p(word)

	// word = strings.ContainsRune("cheese")
	// p(word)

	abc := strings.Count("five", "")
	p(abc)

	p(strings.ToLower("GO GOLANG"))
	p(strings.ToUpper("go python and cplusplus"))

	s = []string{"Orange", "hello", "the new Black"}
	a = strings.Join(s, "so what")
	p(a)

}
