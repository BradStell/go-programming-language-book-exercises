package unitconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string {
	return fmt.Sprintf("%g Feet", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g Meters", m)
}
