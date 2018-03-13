package overlapping_rectangles

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {
	for _, testCase := range []struct {
		rectangle Rectangle
		area      float64
	}{
		{
			rectangle: Rectangle{Point{2, 1}, Point{5, 5}},
			area:      12,
		},
	} {
		area := testCase.rectangle.Area()
		if area != testCase.area {
			t.Errorf("expected error to be %f got %f", testCase.area, area)
		}
	}
}

func TestIntersection(t *testing.T) {
	for _, testCase := range []struct {
		rectangleA  Rectangle
		rectangleB  Rectangle
		rectangleC  Rectangle
		expectError bool
	}{
		{
			rectangleA: Rectangle{Point{2, 1}, Point{5, 5}},
			rectangleB: Rectangle{Point{3, 2}, Point{5, 7}},
			rectangleC: Rectangle{Point{3, 2}, Point{5, 5}},
		},
		{
			rectangleA: Rectangle{Point{0, 1}, Point{3, 3}},
			rectangleB: Rectangle{Point{5, 0}, Point{10, 4}},
			rectangleC: Rectangle{Point{0, 0}, Point{0, 0}},
		},
		{
			rectangleA: Rectangle{Point{7, 2}, Point{9, 3}},
			rectangleB: Rectangle{Point{7, 5}, Point{9, 6}},
			rectangleC: Rectangle{Point{0, 0}, Point{0, 0}},
		},
		{
			rectangleA: Rectangle{Point{7, 2}, Point{9, 4}},
			rectangleB: Rectangle{Point{7, 3}, Point{9, 6}},
			rectangleC: Rectangle{Point{7, 3}, Point{9, 4}},
		},
	} {
		rectangle := testCase.rectangleA.Intersect(testCase.rectangleB)
		if !reflect.DeepEqual(rectangle, testCase.rectangleC) {
			t.Errorf("expected error to be %#v got %#v", testCase.rectangleC, rectangle)
		}
	}
}
