package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	cx := 0.5
	cy := 0.5
	cr := 0.5

	inside := 0
	numSamples := 1000000

	for i := 0; i < numSamples; i++ {
		pointx := rand.Float64()
		pointy := rand.Float64()

		d := math.Sqrt(math.Pow(pointx-cx, 2) + math.Pow(pointy-cy, 2))

		if d < cr {
			inside += 1
		}
	}

	area := float64(inside) / float64(numSamples) // this is pi * r ^ 2
	pi := area / (cr * cr)                        // pi is area / r ^ 2
	fmt.Println(inside, pi)
}
