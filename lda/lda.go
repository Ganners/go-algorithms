package lda

import (
	"math/rand"
)

type Corpus struct {
	wordMatrix  [][]string
	classMatrix [][]int
	numClasses  int
	alpha       float64
	beta        float64
}

func (c *Corpus) GibbsSample() {
	for j := range c.wordMatrix {
		for i := range c.wordMatrix[j] {
			// get probabilities within current row only
			mask := c.FullMask()
			c.UnmaskRow(mask, j)
			c.MaskEntry(mask, j, i)
			localClassesProba := c.Histogram(mask)
			VecDiv(localClassesProba, VecSum(localClassesProba))

			mask = c.FullMask()
			c.UnmaskWord(mask, c.wordMatrix[j][i])
			c.MaskEntry(mask, j, i)
			globalClassesProba := c.Histogram(mask)
			VecDiv(globalClassesProba, VecSum(globalClassesProba))

			VecAdd(localClassesProba, c.alpha)
			VecAdd(globalClassesProba, c.beta)

			classProbas := VecMul(localClassesProba, globalClassesProba)
			// c.classMatrix[j][i] = ArgMax(classProbas)
			c.classMatrix[j][i] = SampleArg(classProbas)
		}
	}
}

func (c *Corpus) FullMask() [][]int {
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

func (c *Corpus) UnmaskRow(mask [][]int, row int) {
	for i := range mask[row] {
		mask[row][i] = 1.
	}
}

func (c *Corpus) UnmaskWord(mask [][]int, searchWord string) {
	for j := range c.wordMatrix {
		for i, word := range c.wordMatrix[j] {
			if word == searchWord {
				mask[j][i] = 1.
			}
		}
	}
}

func (c *Corpus) MaskEntry(mask [][]int, j, i int) {
	mask[j][i] = 0.
}

func (c *Corpus) Histogram(mask [][]int) []float64 {
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

func (c *Corpus) TopDocumentClasses() []int {
	documentClasses := make([]int, len(c.wordMatrix))
	for j := range documentClasses {
		mask := c.FullMask()
		c.UnmaskRow(mask, j)
		hist := c.Histogram(mask)
		documentClasses[j] = ArgMax(hist)
	}
	return documentClasses
}

func NewCorpus(documents []string, numClasses int) *Corpus {
	wordMatrix := make([][]string, len(documents))
	classMatrix := make([][]int, len(documents))
	for j, document := range documents {
		wordMatrix[j] = wordsFromDocument(document)
		docWordClasses := make([]int, len(wordMatrix[j]))
		for i := range docWordClasses {
			docWordClasses[i] = rand.Intn(numClasses)
		}
		classMatrix[j] = docWordClasses
	}
	return &Corpus{
		wordMatrix:  wordMatrix,
		classMatrix: classMatrix,
		numClasses:  numClasses,
		alpha:       1e-1,
		beta:        1e-2,
	}
}
