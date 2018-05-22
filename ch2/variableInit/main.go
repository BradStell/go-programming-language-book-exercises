package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	// normal variabnle declaration (using var or const)
	// var <name> <type> = <expression>
	// either <type> or <expression> may be eliminated (not both)
	var i, j, k int = 1, 2, 3
	var b, f, s = true, 2.3, "four"

	fmt.Printf("i = %x\n", i)
	fmt.Printf("j = %x\n", j)
	fmt.Printf("k = %x\n", k)
	fmt.Printf("b = %t\n", b)
	fmt.Printf("f = %g\n", f)
	fmt.Printf("s = %s\n", s)

	var a, c, d = someFunc()
	fmt.Printf("a = %b\n", a)
	fmt.Printf("c = %b\n", c)
	fmt.Printf("d = %b\n", d)

	// Shorthand variable declaration
	// only permittin for local variables
	// <name> := <expression>
	// type of <name> determined by type of <expression>
	name := "brad"
	age := 30
	pi := 3.14
	fmt.Printf("name = %s\nage = %d\npi = %g\n", name, age, pi)

	// shorthand is the goto
	// var is used when variable will be assigned later
	// or when variable needs different type than the expression
	// ex.
	var boiling float64 = 100
	fmt.Printf(" boiling= %g\n", boiling)

	var names []string
	var err error
	var p Point
	fmt.Printf("names = %v\nerr = %v\np = %v\n", names, err, p)

	// can also declare and assign multiple shorthand variables
	n, m := 0, 1
	fmt.Printf("n = %d\nm = %d\n", n, m)

	// := is a declaration and = is an assignment
	// swap values of a and c
	a, c = c, a

	// if a variable already exists := will only assign it, not redeclare it
	// this declares and assigns ui but only assings a
	ui, a := 756, 45
	fmt.Printf("ui = %d\n", ui)

	// := must at least declare at least 1 new variable or else it wont compile
}

func someFunc() (int, int, int) {
	return 113, 114, 115
}
