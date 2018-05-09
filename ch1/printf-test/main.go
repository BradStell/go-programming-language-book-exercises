package main

import "fmt"

func main() {

	var number int = 1987
	var pi float32 = 3.14

	fmt.Printf("Decimal Integer:\t%d\n", number)

	fmt.Printf("Hexadecimal:\t%x\n", number)

	fmt.Printf("Octal:\t%o\n", number)

	fmt.Printf("Binary:\t%b\n", number)

	fmt.Printf("Floating Point:\t%f\n", pi)

	fmt.Printf("Unicode:\t%c\n", number)

	fmt.Printf("Type:\t%T\n", number)
}
