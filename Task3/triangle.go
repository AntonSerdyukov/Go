// Package triangle contains mathods to detect what kind given triangle is.
//
package triangle

import (
	"math"
)

// Kind triangle type representation
type Kind int

const (
	// NaT is not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
	// Deg degenerate triangle
	Deg
)

func KindFromSides(a float64, b float64, c float64) Kind {
	if math.IsNaN(a) == true || math.IsNaN(b) == true || math.IsNaN(c) == true {
		return NaT
	}
	if a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1) {
		return NaT
	}
	if a == math.Inf(-1) || b == math.Inf(-1) || c == math.Inf(-1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}
	if a+b == c || b+c == a || a+c == b {
		return Deg
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}
