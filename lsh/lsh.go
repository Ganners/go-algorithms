package lsh

import (
	"math"
)

const (
	BucketBucketSize = 64
)

type Thing struct {
	Location Vector
	Data     interface{}
}

type Bucket struct {
	hyperplanes []Vector
	things      [][]*Thing
}

func ComputeHash(location Vector, hyperplanes []Vector) uint64 {
	hash := uint64(0)
	for i, hyperplane := range hyperplanes {
		dist := location.Dot(hyperplane)
		hash |= uint64(math.Float32bits(dist)&(1<<31)>>31) << i
	}
	return hash % BucketBucketSize
}

func NewBucket(things []*Thing, numHyperplanes int) *Bucket {
	d := len(things[0].Location)
	hyperplanes := make([]Vector, numHyperplanes)
	for i := range hyperplanes {
		hyperplanes[i] = NewRandUnitVector(d)
	}
	bucket := &Bucket{
		hyperplanes: hyperplanes,
		things:      make([][]*Thing, BucketBucketSize),
	}
	for _, thing := range things {
		bucket.Add(thing)
	}
	return bucket
}

func (b *Bucket) Add(thing *Thing) {
	location := thing.Location.Norm()
	hash := ComputeHash(location, b.hyperplanes)
	b.things[hash] = append(b.things[hash], thing)
}

func (b *Bucket) Get(location Vector) []*Thing {
	hash := ComputeHash(location, b.hyperplanes)
	return b.things[hash]
}

// LSHIndex is a locality sensitive hashing KNN index, which uses cosine
// similarity as the distance metric (normalization will be handled)
type LSHIndex struct {
	buckets []*Bucket
}

func NewLSHIndex(things []*Thing, numHyperplanes, numBuckets int) *LSHIndex {
	buckets := make([]*Bucket, numBuckets)
	for i := range buckets {
		buckets[i] = NewBucket(things, numHyperplanes)
	}
	return &LSHIndex{
		buckets: buckets,
	}
}

func (l *LSHIndex) TopK(location Vector, k int) []*Thing {
	location = location.Norm()
	closeByConcat := make([]*Thing, 0)
	for _, bucket := range l.buckets {
		closeBy := bucket.Get(location)
		if closeBy != nil {
			closeByConcat = append(closeBy, closeBy...)
		}
	}

	if len(closeByConcat) == 0 {
		return nil
	}

	closeBySet := make(map[*Thing]struct{})
	for _, thing := range closeByConcat {
		closeBySet[thing] = struct{}{}
	}

	closeBy := make([]*Thing, 0, len(closeBySet))
	for thing := range closeBySet {
		closeBy = append(closeBy, thing)
	}

	distances := make(Vector, len(closeBy))
	for i := range distances {
		distances[i] = location.Dot(closeBy[i].Location.Norm())
	}

	_, indices := distances.SortDesc()

	closeBySorted := make([]*Thing, len(indices))
	for i, idx := range indices {
		closeBySorted[i] = closeBy[idx]
	}

	return closeBySorted[:min(k, len(closeBySorted))]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
