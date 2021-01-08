package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
)

func readDocuments(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory: %v", err)
	}

	documents := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fp := filepath.Join(directory, file.Name())
		log.Printf("reading file: %s", fp)
		fileContents, err := ioutil.ReadFile(fp)
		if err != nil {
			return nil, fmt.Errorf("unable to read file %s: %v", file.Name(), err)
		}
		documents = append(documents, string(fileContents))
	}
	return documents, nil
}

func splitCorpus(documents []string) []string {
	corpus := []string{}
	for _, document := range documents {
		document = strings.ToUpper(document)
		document = strings.ReplaceAll(document, `"`, "")
		document = strings.ReplaceAll(document, `(`, "")
		document = strings.ReplaceAll(document, `)`, "")
		corpus = append(corpus, strings.Fields(document)...)
	}
	return corpus
}

type Histogram map[string]float64

func (h Histogram) Normalize() Histogram {
	sumOfTransitions := 0.0
	for _, count := range h {
		sumOfTransitions += count
	}
	for word, count := range h {
		h[word] = count / sumOfTransitions
	}
	return h
}

type TransitionHistogram map[string]Histogram

func (t TransitionHistogram) Normalize() TransitionHistogram {
	for word, hist := range t {
		t[word] = hist.Normalize()
	}
	return t
}

type MarkovChain struct {
	Order               int
	TransitionHistogram TransitionHistogram
}

func NewMarkovChain(order int) *MarkovChain {
	return &MarkovChain{
		Order:               order,
		TransitionHistogram: make(TransitionHistogram),
	}
}

func (m *MarkovChain) Fit(corpus []string) {
	transitionHistogram := make(TransitionHistogram)
	for i := 0; i < len(corpus)-m.Order; i += m.Order {
		head := corpus[i]
		tail := corpus[i+m.Order]
		if _, ok := transitionHistogram[head]; !ok {
			transitionHistogram[head] = make(Histogram)
		}
		transitionHistogram[head][tail] += 1.
	}
	m.TransitionHistogram = transitionHistogram.Normalize()
}

func (m *MarkovChain) Generate(seed []string, numWords int) []string {
	generatedDoc := seed
	for i := 0; i < numWords; i++ {

		// loop from order to 1 (down), pull out each hist from some structure storing order to hist (also need to generate)
		// then average or multiply or something...?
		hist := m.TransitionHistogram[generatedDoc[len(generatedDoc)-m.Order]]
		randPoint, currPoint := rand.Float64(), 0.0
		for nextWord, proba := range hist {
			currPoint += proba
			if currPoint > randPoint {
				generatedDoc = append(generatedDoc, nextWord)
				break
			}
		}
	}
	return generatedDoc
}

func main() {
	documents, err := readDocuments("./data")
	if err != nil {
		log.Fatalf("could not read documents: %v", err)
	}
	mc := NewMarkovChain(2)
	mc.Fit(splitCorpus(documents))
	generated := mc.Generate([]string{"HAUNT", "ME"}, 32)
	log.Println(generated)
}
