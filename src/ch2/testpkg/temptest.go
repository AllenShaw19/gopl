package testpkg

import "ch2/tempconv"

func GetF(c float64) float64 {
	return float64(tempconv.CToF(tempconv.Celsius(c)))
}
