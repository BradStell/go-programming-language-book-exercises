package tempconv

import "fmt"

// Celsius - type float64
type Celsius float64

// Fahrenheit - type float64
type Fahrenheit float64

// AbsoluteZeroC - Celsius value of -273.15
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
