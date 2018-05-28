// Package tempconv performs Celsius and Fahrenheit temperature computations
package tempconv

import "fmt"

// type <name> <underlying-type>
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Types can have "behaviors" associated with them
// in the form of 'type methods'
// (c Celsius) appearing before the function name associates
// with the Celsius type a method named String
// this method will be automatically called when using fmt.Print*
// fmt.Printf("%v", c) -> 100ºC
// fmt.Println(c) -> 100ºC
func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}
