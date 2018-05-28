package unitconv

const feetInMeter = 3.28

func FeetToMeters(f Feet) Meter {
	return Meter(f / Feet(feetInMeter))
}

func MetersToFeet(m Meter) Feet {
	return Feet(m * Meter(feetInMeter))
}
