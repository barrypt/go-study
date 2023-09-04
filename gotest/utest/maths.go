package utest

import "math"

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
    return math.Abs(x)
}

// Max returns the larger of x or y.
func Max(x, y float64) float64 {
    return math.Max(x, y)
}