package lda

import (
	"testing"

	"github.com/go-test/deep"
)

func TestGibbsSampling(t *testing.T) {
	corpus := Corpus2{
		wordMatrix: [][]string{
			{"ball", "ball", "ball", "planet", "galaxy"},
			{"referendum", "planet", "planet", "referendum", "referendum"},
			{"planet", "planet", "galaxy", "planet", "ball"},
			{"planet", "galaxy", "referendum", "planet", "ball"},
		},
		classMatrix: [][]int{
			{1, 0, 0, 2, 2},
			{1, 2, 2, 1, 2},
			{2, 0, 2, 1, 1},
			{2, 0, 1, 2, 0},
		},
		numClasses: 3,
		alpha:      0.01,
		beta:       0.25,
	}

	expectedCorpus := Corpus2{
		wordMatrix: [][]string{
			{"ball", "ball", "ball", "planet", "galaxy"},
			{"referendum", "planet", "planet", "referendum", "referendum"},
			{"planet", "planet", "galaxy", "planet", "ball"},
			{"planet", "galaxy", "referendum", "planet", "ball"},
		},
		classMatrix: [][]int{
			{0, 0, 0, 2, 0},
			{1, 2, 1, 1, 1},
			{2, 2, 2, 2, 0},
			{2, 2, 1, 2, 0},
		},
		numClasses: 3,
		alpha:      0.01,
		beta:       0.05,
	}

	t.Error("expected", expectedCorpus.classMatrix)
	t.Error("actual  ", corpus.classMatrix)

	deep.CompareUnexportedFields = true
	if diff := deep.Equal(corpus, expectedCorpus); diff != nil {
		t.Errorf("corpuses do not match: %v", diff)
	}
}
