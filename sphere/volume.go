package sphere

import (
	"math"
)

func ComputeVolume(radius float64) float64 {
	return (4. / 3.) * (math.Pi * (radius * radius * radius))
}

func ComputeVolumeCubeRatios(radius float64) float64 {
	ratio := 1.8279393035113453702678043555351905524730682373046875
	lower := func(radius float64) float64 { return 2 * radius * radius * radius }
	upper := func(radius float64) float64 { return 2 * math.Pow(math.Sqrt(radius*radius+radius*radius), 3) }
	return (lower(radius) + upper(radius)) / ratio
}
