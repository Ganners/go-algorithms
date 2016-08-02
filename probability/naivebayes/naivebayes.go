package naivebayes

import (
	"errors"
	"fmt"
	"math"
)

// Classifier is what is returned from the classify receiver
type classifier struct {
	featureName string
	cardinality int
	features    map[string][]float64
	params      map[string]parameters
}

// Parameters represent gaussian distribution assumption
type parameters struct {
	mean   float64
	stddev float64
}

// NaiveBayes takes a table, and can classify and make predictions
// based on those classifications
//
// The API needs some work, but it's okay
type NaiveBayes struct {
	table []map[string]float64
}

// NewNaiveBayes creates a pointer to a NaiveBayes
func NewNaiveBayes(trainingData []map[string]float64) *NaiveBayes {
	return &NaiveBayes{
		table: trainingData,
	}
}

// Performs a prediction
func (cl *NaiveBayes) Predict(
	featureClassifier map[float64]*classifier,
	row map[string]float64,
) (float64, float64) {

	total := 0

	for _, classifier := range featureClassifier {
		total += classifier.cardinality
	}

	numerator := math.Inf(-1)
	posteriorNumerator := 0.0

	for posterior, classifier := range featureClassifier {
		probability := float64(classifier.cardinality) / float64(total)
		for name, params := range classifier.params {

			// Check the row exists
			rowValue, ok := row[name]
			if !ok {
				return math.Inf(-1), math.Inf(-1)
			}

			avg := params.mean
			std := params.stddev

			// Determine the distribution for the posterior
			probability *= (1 / math.Sqrt(2*math.Pi*std)) *
				math.Exp((-math.Pow(rowValue-avg, 2))/(2*std))
		}

		// If this has a greater probability than what is there then
		// keep track of it
		if probability > numerator {
			numerator = probability
			posteriorNumerator = posterior
		}
	}

	// Return the classification and posterior
	return posteriorNumerator, numerator
}

// Generates a classification map which can be used for predictions
func (nb *NaiveBayes) Classify(fieldName string) (error, map[float64]*classifier) {

	if len(nb.table) == 0 {
		return errors.New("No data to classify"), nil
	}

	// Look for presence of the field name
	for _, row := range nb.table {
		if _, found := row[fieldName]; !found {
			return fmt.Errorf("Could not find %s in all rows", fieldName), nil
		}
	}

	dataByFeature := make(map[float64]*classifier, 10)

	for _, row := range nb.table {

		featureVal := row[fieldName]

		// Build inner map if it is nil
		if dataByFeature[featureVal] == nil {
			dataByFeature[featureVal] = &classifier{
				params:   make(map[string]parameters, 10),
				features: make(map[string][]float64, 10),
			}
		}

		// Incremement number in this set
		dataByFeature[featureVal].cardinality++

		featureParams := dataByFeature[featureVal].features

		for field, value := range row {
			if field == fieldName {
				continue
			}
			featureParams[field] = append(featureParams[field], value)
		}
	}

	// Build gaussian parameters table
	for _, classification := range dataByFeature {
		for feature, features := range classification.features {
			classification.params[feature] = parameters{
				mean:   mean(features),
				stddev: stddev(features),
			}
		}
	}

	return nil, dataByFeature
}

// Calculate the mean
func mean(features []float64) float64 {
	sum := 0.0
	for _, f := range features {
		sum += f
	}
	return sum / float64(len(features))
}

// Calculate the standard deviation from the mean
func stddev(features []float64) float64 {
	avg := mean(features)
	n := float64(len(features) - 1)
	sum := 0.0
	for _, f := range features {
		sum += math.Pow(f-avg, 2)
	}
	return sum / n
}
