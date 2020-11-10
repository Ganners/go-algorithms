package cosdist

import (
	"math"
)

// CosineDistance computes what it says on the tin, but all in 32 bits
func CosineDistance(a, b []float32) float32 {
	// if the vector lengths don't line up then return a negative score
	if len(a) != len(b) {
		return -1
	}

	// Compute all dot products
	// ab, aa, bb := 0.0, 0.0, 0.0
	var ab, aa, bb float32
	for i := range a {
		ab += a[i] * b[i]
		aa += a[i] * a[i]
		bb += b[i] * b[i]
	}

	return ab / float32(math.Sqrt(float64(aa*bb)))
}

func CosineDistanceAVX(a, b []float32) float32

// Dot computes what it says on the tin, but all in 32 bits
func Dot(a, b []float32) float32 {
	// if the vector lengths don't line up then return a negative score
	if len(a) != len(b) {
		return -1
	}

	// Compute all dot products
	var ab float32
	for i := range a {
		ab += a[i] * b[i]
	}

	return ab
}

func DotAVX(a, b []float32) float32
