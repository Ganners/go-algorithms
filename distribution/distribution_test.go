package distribution

import (
	"strconv"
	"testing"
)

// Basic test to see if we fill 10 buckets with 20 runs
func TestRandomDistribution(t *testing.T) {
	buckets := make([]int, 10)
	for i := 0; i < 20; i++ {

		dist := RandomDistribution("Something Constant", "User-"+strconv.Itoa(i))
		dist2 := RandomDistribution("Something Constant", "User-"+strconv.Itoa(i))

		bucket := int(dist * 10.0)
		buckets[bucket]++

		// Dist and dist2 should be the same
		if dist != dist2 {
			t.Errorf("Different distributions calculated for the same input")
		}
	}
	// Out of 100, each should have at least 1
	for b, count := range buckets {
		if count == 0 {
			t.Errorf("Expected to receive count greater than 0 for bucket %d", b)
		}
	}
}
