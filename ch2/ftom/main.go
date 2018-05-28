package main

import (
	"fmt"
	"go-book-examples/ch2/unitconv"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	f := unitconv.Feet(t)
	m := unitconv.Meter(t)

	fmt.Printf("%v == %v\n", f, unitconv.FeetToMeters(f))
	fmt.Printf("%v == %v\n", m, unitconv.MetersToFeet(m))
}
