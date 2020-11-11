package lda

import (
	"math/rand"
)

type Corpus2 struct {
	wordMatrix  [][]string
	classMatrix [][]int
	numClasses  int
	alpha       float64
	beta        float64
}

func (c *Corpus2) GibbsSample() {
	for j := range c.wordMatrix {
		for i := range c.wordMatrix[j] {
			// get probabilities within current row only
			localMask := c.FullMask()
			c.UnmaskRow(localMask, j)
			c.MaskEntry(localMask, j, i)

			localClassesProba := c.Histogram(localMask)
			VecDiv(localClassesProba, VecSum(localClassesProba))

			globalMask := c.FullMask()
			c.UnmaskWord(globalMask, c.wordMatrix[j][i])
			c.MaskEntry(globalMask, j, i)
			globalClassesProba := c.Histogram(globalMask)
			VecDiv(globalClassesProba, VecSum(globalClassesProba))

			VecAdd(localClassesProba, c.alpha)
			VecAdd(globalClassesProba, c.beta)

			classProbas := VecMul(localClassesProba, globalClassesProba)

			c.classMatrix[j][i] = ArgMax(classProbas)
		}
	}
}

func (c *Corpus2) FullMask() [][]int {
	mask := make([][]int, len(c.classMatrix))
	for j := range mask {
		rowMask := make([]int, len(c.classMatrix[j]))
		for i := range rowMask {
			rowMask[i] = 0.
		}
		mask[j] = rowMask
	}
	return mask
}

func (c *Corpus2) UnmaskRow(mask [][]int, row int) {
	for i := range mask[row] {
		mask[row][i] = 1.
	}
}

func (c *Corpus2) UnmaskWord(mask [][]int, searchWord string) {
	for j := range c.wordMatrix {
		for i, word := range c.wordMatrix[j] {
			if word == searchWord {
				mask[j][i] = 1.
			}
		}
	}
}

func (c *Corpus2) MaskEntry(mask [][]int, j, i int) {
	mask[j][i] = 0.
}

func (c *Corpus2) Histogram(mask [][]int) []float64 {
	hist := make([]float64, c.numClasses)
	for j := range c.classMatrix {
		for i := range c.classMatrix[j] {
			if mask[j][i] == 0 {
				continue
			}
			hist[c.classMatrix[j][i]] += 1.
		}
	}
	return hist
}

func NewCorpus(documents []string, numClasses int) *Corpus2 {
	wordMatrix := make([][]string, len(documents))
	classMatrix := make([][]int, len(documents))
	for i, document := range documents {
		wordMatrix[i] = wordsFromDocument(document)
		docWordClasses := make([]int, len(wordMatrix[i]))
		for i := range docWordClasses {
			docWordClasses[i] = rand.Intn(numClasses)
		}
		classMatrix[i] = docWordClasses
	}
	return &Corpus2{
		wordMatrix:  wordMatrix,
		classMatrix: classMatrix,
		alpha:       1e-10,
		beta:        1e-10,
	}
}
