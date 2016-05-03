package atoi

import "testing"

func TestParseFloat(t *testing.T) {

	testCases := []struct {
		In        string
		ExpectVal float64
		ExpectErr error
	}{
		{
			In:        "11",
			ExpectVal: 11,
			ExpectErr: nil,
		},
		{
			In:        "1.1",
			ExpectVal: 1.1,
			ExpectErr: nil,
		},
		{
			In:        "3.1415",
			ExpectVal: 3.1415,
			ExpectErr: nil,
		},
		{
			In:        "2.71828",
			ExpectVal: 2.71828,
			ExpectErr: nil,
		},
		{
			In:        "-2.71828",
			ExpectVal: -2.71828,
			ExpectErr: nil,
		},
		{
			In:        "-2.718281.8284",
			ExpectVal: 0,
			ExpectErr: StringNotValid,
		},
	}

	for _, test := range testCases {

		out, err := parseFloat(test.In)

		if out != test.ExpectVal {
			t.Errorf("Expected output %f, got %f", test.ExpectVal, out)
		}
		if err != test.ExpectErr {
			t.Errorf("Expected error %v, got %v", test.ExpectErr, err)
		}
	}
}
