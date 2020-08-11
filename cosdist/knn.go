package cosdist

import "sort"

type Index struct {
	EntryIDs     []string
	EntryVectors [][]float32
}

func NewIndex(entryIDs []string, entryVectors [][]float32) *Index {
	return &Index{
		EntryIDs:     entryIDs,
		EntryVectors: entryVectors,
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

func (idx *Index) computeDistances(centroid []float32) Distances {
	distances := make([]distanceTuple, len(idx.EntryIDs))
	for i := range idx.EntryIDs {
		distances[i] = distanceTuple{
			id:       idx.EntryIDs[i],
			distance: CosineDistanceAVX(centroid, idx.EntryVectors[i]),
		}
	}
	return distances
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
