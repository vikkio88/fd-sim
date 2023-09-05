package utils

import "math"

func ExpFactor(value, target, result float64) float64 {
	return result + (value-target)*(1-math.Pow(2, -math.Abs(value-target)))
}
