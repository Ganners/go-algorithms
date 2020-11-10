package lsh

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestLSH(t *testing.T) {
	things := []*Thing{
		{Location: Vector([]float32{.9, .9, .9})},
		{Location: Vector([]float32{.9, .8, .8})},
		{Location: Vector([]float32{.9, .7, .7})},
		{Location: Vector([]float32{.4, .1, .1})},
		{Location: Vector([]float32{.3, .2, .1})},
		{Location: Vector([]float32{-.3, -.2, -.1})},
	}

	index := NewLSHIndex(things, 5, 5)
	expectedTopK := []*Thing{
		{Location: Vector([]float32{-.3, -.2, -.1})},
	}
	topK := index.TopK(Vector([]float32{-0.33, -0.21, -0.1}), 1)
	if !reflect.DeepEqual(topK, expectedTopK) {
		t.Errorf("topK %#v does not match expected %#v", topK, expectedTopK)
	}

	topK = index.TopK(Vector([]float32{0.33, 0.21, 0.1}), 1)
	expectedTopK = []*Thing{
		{Location: Vector([]float32{.3, .2, .1})},
	}
	if !reflect.DeepEqual(topK, expectedTopK) {
		t.Errorf("topK %#v does not match expected %#v", topK, expectedTopK)
	}
}

func generateEmbedding(d int) Vector {
	vec := make([]float32, d)
	for i := 0; i < d; i++ {
		vec[i] = float32(rand.NormFloat64())
	}
	return vec
}

func generateIndex(n, d int) *LSHIndex {
	things := make([]*Thing, n)
	for i := 0; i < n; i++ {
		things[i] = &Thing{Location: generateEmbedding(d)}
	}
	return NewLSHIndex(things, 100, 25)
}

func BenchmarkTopK_10000_768_50(b *testing.B) {
	n, d, k := 10000, 768, 50
	index := generateIndex(n, d)
	centroid := generateEmbedding(d)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index.TopK(centroid, k)
	}
}

func BenchmarkTopK_100000_768_50(b *testing.B) {
	n, d, k := 100000, 768, 50
	index := generateIndex(n, d)
	centroid := generateEmbedding(d)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index.TopK(centroid, k)
	}
}
