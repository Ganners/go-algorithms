package main

import (
	"log"
	"sort"
)

// Rating represents a user -> item mapping with a float rating
type Rating struct {
	User   string
	Item   string
	Rating float64
}

type Matrix struct {
	XLabels []string
	YLabels []string

	matrix [][]float64
}

type Ratings []Rating

func (r Ratings) ToMatrix() Matrix {
	// construct unique X +Y labels
	xLabels := []string{}
	yLabels := []string{}

	xLabelSet := map[string]struct{}{}
	yLabelSet := map[string]struct{}{}
	zRatingSet := map[string]float64{}

	for _, entry := range r {
		xLabelSet[entry.User] = struct{}{}
		yLabelSet[entry.Item] = struct{}{}

		zRatingSet[entry.User+entry.Item] = entry.Rating
	}

	for label := range xLabelSet {
		xLabels = append(xLabels, label)
	}
	for label := range yLabelSet {
		yLabels = append(yLabels, label)
	}

	sort.Strings(xLabels)
	sort.Strings(yLabels)

	matrix := make([][]float64, len(xLabels))
	for u := range matrix {
		matrix[u] = make([]float64, len(yLabels))
	}
	for u, xLabel := range xLabels {
		for i, yLabel := range yLabels {
			rating, _ := zRatingSet[xLabel+yLabel]
			matrix[u][i] = rating
		}
	}
	return Matrix{
		XLabels: xLabels,
		YLabels: yLabels,

		matrix: matrix,
	}
}

func getRatings() Ratings {
	return Ratings{
		// Jack
		{
			User:   "Jack",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Jack",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Jack",
			Item:   "Milk",
			Rating: 1.,
		},

		// Mary
		{
			User:   "Mary",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Mary",
			Item:   "Milk",
			Rating: 1.,
		},
		{
			User:   "Mary",
			Item:   "Beef",
			Rating: 1.,
		},

		// Jane
		{
			User:   "Jane",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Jane",
			Item:   "Butter",
			Rating: 1.,
		},

		// Sayani
		{
			User:   "Sayani",
			Item:   "Bread",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Milk",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Beef",
			Rating: 1.,
		},
		{
			User:   "Sayani",
			Item:   "Ham",
			Rating: 1.,
		},

		// John
		{
			User:   "John",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "John",
			Item:   "Ham",
			Rating: 1.,
		},

		// Tom
		{
			User:   "Tom",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Tom",
			Item:   "Beef",
			Rating: 1.,
		},
		{
			User:   "Tom",
			Item:   "Ham",
			Rating: 1.,
		},

		// Peter
		{
			User:   "Peter",
			Item:   "Butter",
			Rating: 1.,
		},
		{
			User:   "Peter",
			Item:   "Fish",
			Rating: 1.,
		},
		{
			User:   "Peter",
			Item:   "Beef",
			Rating: 1.,
		},

		// OutOfSample
		{
			User:   "OutOfSample",
			Item:   "Eels",
			Rating: 1.,
		},
	}
}

func main() {
	ratings := getRatings()
	matrix := ratings.ToMatrix()

	log.Println(matrix)
}
