package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func loadData() ([][]float64, []float64, error) {
	X := make([][]float64, 0, 32)
	y := make([]float64, 0, 32)

	filename := "data/winequality-white.csv"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return X, y, err
	}

	line := 0
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ';'

	for {
		line++
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return X, y, err
		}

		// skip header
		if line == 1 {
			continue
		}

		if len(record) != 12 {
			return X, y, errors.New("expecting 12 columns")
		}

		R := make([]float64, len(record))
		for i, r := range record {
			R[i], err = strconv.ParseFloat(r, 64)
			if err != nil {
				return X, y, err
			}
		}

		X = append(X, R[:len(R)-1])
		y = append(y, R[len(R)-1])
	}

	return X, y, nil
}

func transpose(slice [][]float64) [][]float64 {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]float64, xl)
	for i := range result {
		result[i] = make([]float64, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// normalize will scale each column vector in-line
func normalize(v [][]float64) [][]float64 {
	for j := 0; j < len(v[0]); j++ {
		min := math.Inf(1)
		max := math.Inf(-1)

		// min-max loop
		for i := range v {
			if v[i][j] < min {
				min = v[i][j]
			}
			if v[i][j] > max {
				max = v[i][j]
			}
		}

		// scale loop
		for i := range v {
			v[i][j] = (v[i][j] - min) / (max - min)
		}
	}
	return v
}

func predict(row []float64, coefficients []float64) float64 {
	ŷ := coefficients[0]
	for i := range row {
		ŷ += coefficients[i+1] * row[i]
	}
	return ŷ
}

func shuffle(
	X [][]float64,
	y []float64,
) (
	[][]float64,
	[]float64,
) {
	p := rand.Perm(len(X))
	for i, j := range p {
		X[j], X[i] = X[i], X[j]
		y[j], y[i] = y[i], y[j]
	}
	return X, y
}

func linearRegressionSGD(
	X [][]float64,
	y []float64,
	learningRate float64,
	numEpochs int,
) []float64 {
	// initialize coefs to 0.0 (with bias coef at start)
	coef := make([]float64, 1+len(X[0]))

	// initialize randomly
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range coef {
		coef[i] = r.Float64()
	}
	// shuffle (to randomize)
	X, y = shuffle(X, y)
	for epoch := 0; epoch < numEpochs; epoch++ {
		sumErr := 0.0
		for i, x := range X {
			ŷ := predict(x, coef)
			err := ŷ - y[i]
			sumErr += err * err
			coef[0] -= learningRate * err

			for j := 1; j < len(coef); j++ {
				coef[j] -= learningRate * err * x[j-1]
			}
		}
	}
	return coef
}

func evaluate(X [][]float64, y []float64, coef []float64) float64 {
	sumErr := 0.0
	for i := range X {
		ŷ := predict(X[i], coef)
		err := ŷ - y[i]
		sumErr += err * err
		// log.Printf("Expected=%.3f, Predicted=%.3f\n", y[i], ŷ)
	}
	return math.Sqrt(sumErr / float64(len(X)))
}

func wine() {
	X, y, err := loadData()
	if err != nil {
		log.Fatal(err)
	}

	// normalize X and y
	X = normalize(X)
	y = transpose(normalize(transpose([][]float64{y})))[0]

	// split test/train
	split := int(float64(len(X)) * 0.8)
	Xtrain, Xtest, ytrain, ytest := X[:split], X[split+1:], y[:split], y[split+1:]

	// compute coefficients from SGD
	best := math.Inf(1)
	bestCoefs := []float64{}
	for _, a := range []float64{0.1, 0.01, 0.001, 0.0001, 0.0005, 0.00001} {
		for _, b := range []int{10, 20, 50, 100, 200, 500, 1000} {
			coef := linearRegressionSGD(Xtrain, ytrain, a, b)
			rmse := evaluate(Xtest, ytest, coef)
			if rmse < best {
				best = rmse
				bestCoefs = coef
			}
		}
	}
	log.Println("best rmse", best, bestCoefs)
}

func ctof() {
	X := [][]float64{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
		{10},
		{20},
		{30},
		{40},
		{50},
		{60},
		{70},
		{80},
		{90},
		{100},
		{200},
		{300},
		{400},
		{500},
		{600},
		{700},
		{800},
		{900},
		{1000},
	}
	y := []float64{
		32.0,
		33.8,
		35.6,
		37.4,
		39.2,
		41.0,
		42.8,
		44.6,
		46.4,
		48.2,
		50.0,
		68.0,
		86.0,
		104.0,
		122.0,
		140.0,
		158.0,
		176.0,
		194.0,
		212.0,
		392.0,
		572.0,
		752.0,
		932.0,
		1112.0,
		1292.0,
		1472.0,
		1652.0,
		1832.0,
	}

	// normalize X and y
	X = normalize(X)
	// y = transpose(normalize(transpose([][]float64{y})))[0]

	// split test/train
	split := int(float64(len(X)) * 0.9)
	Xtrain, Xtest, ytrain, ytest := X[:split], X[split+1:], y[:split], y[split+1:]

	// compute coefficients from SGD
	coef := linearRegressionSGD(Xtrain, ytrain, 0.1, 100)
	rmse := evaluate(Xtest, ytest, coef)
	log.Println("rmse", coef, rmse, predict([]float64{20 / 1000}, coef))
}

func main() {
	wine()
	ctof()
}
