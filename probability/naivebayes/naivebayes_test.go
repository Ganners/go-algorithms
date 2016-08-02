package naivebayes

import "testing"

// Input table
var data = []map[string]float64{
	{
		// Tall fat outlier
		"gender":   0.0,
		"height":   9.0,
		"weight":   380.0,
		"footSize": 10,
	},
	{
		"gender":   0.0,
		"height":   6.0,
		"weight":   180.0,
		"footSize": 12,
	},
	{
		"gender":   0.0,
		"height":   5.92,
		"weight":   190.0,
		"footSize": 11,
	},
	{
		"gender":   0.0,
		"height":   5.58,
		"weight":   170.0,
		"footSize": 12,
	},
	{
		"gender":   0.0,
		"height":   5.92,
		"weight":   165.0,
		"footSize": 10,
	},
	{
		"gender":   1.0,
		"height":   5.0,
		"weight":   100.0,
		"footSize": 6,
	},
	{
		"gender":   1.0,
		"height":   5.5,
		"weight":   150.0,
		"footSize": 8,
	},
	{
		"gender":   1.0,
		"height":   5.42,
		"weight":   130.0,
		"footSize": 7,
	},
	{
		"gender":   1.0,
		"height":   5.75,
		"weight":   150.0,
		"footSize": 9,
	},
}

func TestGenderClassifier(t *testing.T) {

	cl := NewNaiveBayes(data)
	err, classifier := cl.Classify("gender")
	if err != nil {
		t.Errorf("Did not expect error, got %s", err)
	}

	for _, test := range []struct {
		features       map[string]float64
		classification float64
	}{
		{
			features: map[string]float64{
				"height":   6,
				"weight":   130.0,
				"footSize": 8,
			},
			classification: 1.0,
		},
		{
			features: map[string]float64{
				"height":   6.0,
				"weight":   180.0,
				"footSize": 12,
			},
			classification: 0.0,
		},
		{
			features: map[string]float64{
				"height":   9.0,
				"weight":   190.0,
				"footSize": 9,
			},
			classification: 0.0,
		},
	} {
		prediction, _ := cl.Predict(classifier, test.features)
		if prediction != test.classification {
			t.Errorf("Prediction %f does not match expected %f", prediction, test.classification)
		}
	}
}

func TestMean(t *testing.T) {
	for _, test := range []struct {
		input  []float64
		output float64
	}{
		{
			input:  []float64{2, 4, 4, 4, 5, 7, 9},
			output: 5,
		},
		{
			input:  []float64{1, 1, 1, 1, 1, 1, 1},
			output: 1,
		},
	} {
		out := mean(test.input)
		if out != test.output {
			t.Errorf("Expected mean to be %f got %f", test.output, out)
		}
	}
}

func TestStddev(t *testing.T) {
	for _, test := range []struct {
		input  []float64
		output float64
	}{
		{
			input:  []float64{2, 4, 4, 4, 5, 7, 9},
			output: 5.33333333333333303727,
		},
		{
			input:  []float64{180, 190, 170, 165},
			output: 122.91666666666667140362,
		},
	} {
		out := stddev(test.input)
		if out != test.output {
			t.Errorf("Expected mean to be %.20f got %.20f", test.output, out)
		}
	}
}
