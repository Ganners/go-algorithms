package lda

import (
	"math/rand"
	"strings"
)

type Document struct {
	words       []string
	wordClasses []int
	classes     int
}

func (d *Document) AssignRandomClasses() {
	d.wordClasses = make([]int, len(d.words))
	for i := range d.wordClasses {
		d.wordClasses[i] = rand.Intn(d.classes)
	}
}

func (d *Document) WordClassCounts(searchWord string) []float64 {
	wordClassCounts := make([]float64, d.classes)
	for i, word := range d.words {
		if word == searchWord {
			wordClassCounts[d.wordClasses[i]] += 1.
		}
	}
	return wordClassCounts
}

func (d *Document) WordClassProbabilities(searchWord string) []float64 {
	filteredWordClasses := make([]int, 0)
	for i, word := range d.words {
		if word == searchWord {
			filteredWordClasses = append(filteredWordClasses, d.wordClasses[i])
		}
	}
	return d.classProbabilities(filteredWordClasses)
}

func (d *Document) GetClass() int {
	classProba := d.AllClassProbabilities()
	return ArgMax(classProba)
}

func (d *Document) AllClassProbabilities() []float64 {
	return d.classProbabilities(d.wordClasses)
}

func (d *Document) classProbabilities(wordClasses []int) []float64 {
	probabilities := make([]float64, d.classes)
	inc := 1. / float64(len(wordClasses))
	for _, class := range wordClasses {
		probabilities[class] += inc
	}
	return probabilities
}

func wordsFromDocument(document string) []string {
	document = strings.ToLower(document)
	document = strings.ReplaceAll(document, "\n", " ")
	wordsRaw := strings.Split(document, " ")
	words := make([]string, 0, len(wordsRaw))
	for _, word := range wordsRaw {
		if word == "" || word == " " {
			continue
		}
		word = strings.TrimSpace(word)
		words = append(words, word)
	}

	return words
}

type Corpus struct {
	documents []*Document
	classes   int
	alpha     float64
	beta      float64
}

func (c Corpus) WordClassProbabilities(searchWord string) []float64 {
	wordClassCounts := make([]float64, c.classes)
	for _, document := range c.documents {
		counts := document.WordClassCounts(searchWord)
		for i := range wordClassCounts {
			wordClassCounts[i] += counts[i]
		}
	}
	VecDiv(wordClassCounts, VecSum(wordClassCounts))
	return wordClassCounts
}

func (c Corpus) GibbsSampling() {
	for _, document := range c.documents {
		for j, word := range document.words {
			// local probabilities for all words
			localProbas := document.AllClassProbabilities()
			outsideWordProbas := c.WordClassProbabilities(word)

			VecAdd(localProbas, c.alpha)
			VecAdd(outsideWordProbas, c.beta)
			classProbas := VecMul(localProbas, outsideWordProbas)

			// assign the best class
			document.wordClasses[j] = ArgMax(classProbas)
		}
	}
}

func LDA(rawDocuments []string, classes, iterations int) []int {
	documents := make([]*Document, len(rawDocuments))
	for i := range rawDocuments {
		document := &Document{
			words:   wordsFromDocument(rawDocuments[i]),
			classes: classes,
		}
		document.AssignRandomClasses()
		documents[i] = document
	}

	corpus := Corpus{
		documents: documents,
		classes:   classes,
		alpha:     1e-10,
		beta:      1e-10,
	}

	for i := 0; i < iterations; i++ {
		corpus.GibbsSampling()
	}

	docClasses := make([]int, len(corpus.documents))
	for i, document := range corpus.documents {
		docClasses[i] = document.GetClass()
	}

	return docClasses
}
