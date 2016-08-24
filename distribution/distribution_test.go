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
		bucket := int(dist * 10.0)
		buckets[bucket]++
	}
	// Out of 100, each should have at least 1
	for b, count := range buckets {
		if count == 0 {
			t.Errorf("Expected to receive count greater than 0 for bucket %d", b)
		}
	}
}
