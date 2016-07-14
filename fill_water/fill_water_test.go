package fill_water

import "testing"

func TestFillWaterCount(t *testing.T) {

	testCases := []struct {
		Environment   []int
		ExpectedCount int
	}{
		{
			Environment:   []int{0, 0},
			ExpectedCount: 0,
		},
		{
			//
			//
			// x
			// x
			// x w x
			// x w x w x
			Environment: []int{
				4, 0, 2, 0, 1,
			},
			ExpectedCount: 3,
		},
		{
			//
			//
			// x
			// x w x
			// x w x
			// x x x w x x x
			Environment: []int{
				4, 1, 3, 0, 1, 1, 1,
			},
			ExpectedCount: 3,
		},
		{
			//
			//
			// x
			// x w x
			// x w x w x
			// x w x w x
			Environment: []int{
				4, 0, 3, 0, 2,
			},
			ExpectedCount: 5,
		},
		{
			//
			//
			// x w w w x
			// x w x w x
			// x w x w x
			// x w x w x
			Environment: []int{
				4, 0, 3, 0, 4,
			},
			ExpectedCount: 9,
		},
		{
			//
			//
			// x
			// x
			// x w w w x
			// x w w w x
			Environment: []int{
				4, 0, 0, 0, 2,
			},
			ExpectedCount: 6,
		},
		{
			//               x
			//       x w w w x
			// x w w x w x w x
			// x w w x w x w x
			// x w w x w x w x
			// x w w x w x w x
			Environment: []int{
				4, 0, 0, 5, 0, 4, 0, 6,
			},
			ExpectedCount: 19,
		},
	}

	for _, test := range testCases {
		res := FillWaterCount(test.Environment)
		if res != test.ExpectedCount {
			t.Errorf("Expected count of %d, got %d", test.ExpectedCount, res)
		}
	}
}
