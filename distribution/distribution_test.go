package distribution

import (
	"strconv"
	"testing"
)

func TestSampleUniformHash(t *testing.T) {
	buckets := make([]int, 10)
	for i := 0; i < 20; i++ {
		dist := SampleUniformHash("Something Constant", "User-"+strconv.Itoa(i))
		dist2 := SampleUniformHash("Something Constant", "User-"+strconv.Itoa(i))
		bucket := int(dist * 10.0)
		buckets[bucket]++

		if dist != dist2 {
			t.Errorf("Different distributions calculated for the same input")
		}
	}
	// each bucket should have at least 1
	for b, count := range buckets {
		if count == 0 {
			t.Errorf("Expected to receive count greater than 0 for bucket %d", b)
		}
	}
}
