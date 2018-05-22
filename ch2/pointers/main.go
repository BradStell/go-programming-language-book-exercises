package main

import "fmt"

func main() {
	// shorthand - variable type inferred
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(*p)
	*p = 2 // equivalent to x = 2
	fmt.Println(x)

	// long hand - variable types explicitly expressed
	var y int = 13
	var pp *int = &y
	fmt.Println(*pp)
	*pp = 23
	fmt.Println(y)

	fmt.Printf("pp address = %v\n", pp)
	fmt.Printf("pp address = %x\n", pp)

	// it is perfectly safe to have a function return the address
	// of a local variable
	bb := f()
	fmt.Printf("bb = %v\nbb = %d\n", bb, *bb)

	// can pass variable address into func to manipulate it
	v := 1
	fmt.Printf("v = %d\n", incr(&v)) // v is now 2
}

func incr(p *int) int {
	*p++
	return *p
}

func f() *int {
	var v int = 13
	return &v
}
