package main

import (
	"fmt"
)

func main() {
	p := new(int) // p, of type *int points to an unnamed int variable
	fmt.Println(*p)
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p)
}
