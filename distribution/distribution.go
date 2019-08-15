package distribution

import (
	"encoding/binary"
	"hash/fnv"
)

// SampleUniformHash returns a good uniform distribution for incrementing
// variables given a constant ID. The value returned will be between 0.0 and 1.0
//
// The first input should be a constant input, perhaps an experiment ID
// The second input should be a variable input, perhaps a user ID
func SampleUniformHash(constant string, variable string) float64 {
	fnvHash := fnv.New64a()
	fnvHash.Write([]byte(constant + variable))
	littleEndian := binary.LittleEndian.Uint64(fnvHash.Sum(nil))
	return float64(littleEndian) / float64(1<<64-1)
}
