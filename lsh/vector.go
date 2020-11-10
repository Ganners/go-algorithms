package lsh

import (
	"math"
	"math/rand"
	"sort"
)

type Vector []float32

func NewRandUnitVector(dims int) Vector {
	vec := make(Vector, dims)
	vec = vec.SeedRandom()
	vec = vec.Divide(vec.Length())
	return vec
}

func (v Vector) SeedRandom() Vector {
	newV := make(Vector, len(v))
	for i := range newV {
		newV[i] = float32(rand.NormFloat64())
	}
	return newV
}

func (v Vector) Divide(denom float32) Vector {
	newV := make(Vector, len(v))
	for i := range newV {
		newV[i] = v[i] / denom
	}
	return newV
}

func (v Vector) Multiply(v2 Vector) Vector {
	newV := make(Vector, len(v))
	for i := range newV {
		newV[i] = v[i] * v2[i]
	}
	return newV
}

func (v Vector) Sum() float32 {
	sum := float32(0.0)
	for _, val := range v {
		sum += val
	}
	return sum
}

func (v Vector) Dot(v2 Vector) float32 {
	newV := make(Vector, len(v))
	for i := range newV {
		newV[i] = v[i] * v2[i]
	}
	return newV.Sum()
}

func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector) Norm() Vector {
	return v.Divide(v.Length())
}

type VectorTuple struct {
	idx   int
	value float32
}

type argSorter []VectorTuple

func (v argSorter) Less(i, j int) bool { return v[i].value < v[j].value }
func (v argSorter) Len() int           { return len(v) }
func (v argSorter) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

func (v Vector) SortDesc() (Vector, []int) {
	tuples := make([]VectorTuple, len(v))
	for i, val := range v {
		tuples[i] = VectorTuple{idx: i, value: val}
	}

	sort.Sort(sort.Reverse(argSorter(tuples)))

	argsorted := make([]int, len(tuples))
	sorted := make(Vector, len(tuples))

	for i, tuple := range tuples {
		argsorted[i] = tuple.idx
		sorted[i] = tuple.value
	}

	return sorted, argsorted
}
