// iota keywords
// iota set the first numerator to zero
// and it increments it by 1

package main

import "fmt"

func main() {
	const (
		c1 = iota // 0
		c2 = iota // 1
		c3 = iota // 2
	)

	fmt.Println(c1, c2, c3)
}
