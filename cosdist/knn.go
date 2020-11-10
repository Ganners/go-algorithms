package cosdist

import (
	"math"
	"sort"
)

type Index struct {
	entryIDs     []string
	entryVectors [][]float32
	knnBatchSize int
}

func NewIndex(entryIDs []string, entryVectors [][]float32) *Index {
	return &Index{
		entryIDs:     entryIDs,
		entryVectors: entryVectors,
		knnBatchSize: 512,
	}
}

type distanceTuple struct {
	id       string
	distance float32
}

type Distances []distanceTuple

func (d Distances) Len() int      { return len(d) }
func (d Distances) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

type ByDistance struct{ Distances }

func (d ByDistance) Less(i, j int) bool {
	return d.Distances[i].distance < d.Distances[j].distance
}

func (idx *Index) computeDistancesBatch(centroid []float32, start, end int) Distances {
	distances := make([]distanceTuple, end-start)
	for i := start; i < end; i++ {
		distances[i-start] = distanceTuple{
			id:       idx.entryIDs[i],
			distance: DotAVX(centroid, idx.entryVectors[i]),
		}
	}
	return distances
}

func (idx *Index) computeDistancesConcurrent(centroid []float32, batchSize int) Distances {
	distances := Distances{}
	distancesCh := make(chan Distances)
	n := len(idx.entryVectors)
	nBatches := int(math.Ceil(float64(n) / float64(batchSize)))

	for start := 0; start < n; start += batchSize {
		end := min(start+batchSize, n)
		go func(centroid []float32, start, end int) {
			distancesCh <- idx.computeDistancesBatch(centroid, start, end)
		}(centroid, start, end)
	}

	for i := 0; i < nBatches; i++ {
		distances = append(distances, []distanceTuple(<-distancesCh)...)
	}

	return distances
}

func (idx *Index) computeDistances(centroid []float32) Distances {
	return idx.computeDistancesConcurrent(centroid, idx.knnBatchSize)
}

func (idx *Index) KNN(centroid []float32, k int) Distances {
	distances := idx.computeDistances(centroid)
	k = min(k, len(distances))
	n := len(distances)
	for i := 0; i < k; i++ {
		maxIndex := i
		maxValue := distances[i].distance
		for j := i + 1; j < n; j++ {
			if distances[j].distance > maxValue {
				maxIndex = j
				maxValue = distances[j].distance
				distances[i], distances[maxIndex] = distances[maxIndex], distances[i]
			}
		}
	}

	distances = distances[:k]
	sort.Sort(sort.Reverse(ByDistance{distances}))
	return distances
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
