package integration

import "testing"

func TestSimpson1(t *testing.T) {

	f := func(x float64) float64 {
		return 1.0
	}

	{
		res := Simpson(f, 0.0, 10.0, 1.0)
		expected := 10.0

		if res != expected {
			t.Errorf("Result %f does not match expected %f", res, expected)
		}
	}
}

func TestSimpsonCubed(t *testing.T) {

	f := func(x float64) float64 {
		return x * x * x
	}

	{
		res := Simpson(f, 0.0, 10.0, 2.0)
		expected := 2500.0

		if res != expected {
			t.Errorf("Result %f does not match expected %f", res, expected)
		}
	}

	{
		res := int(Simpson(f, 0.0, 10.0, 100000.0))
		expected := 2500

		if res != expected {
			t.Errorf("Result %d does not match expected %d", res, expected)
		}
	}
}

func TestSimpsonQuad(t *testing.T) {

	f := func(x float64) float64 {
		return x * x * x * x
	}

	{
		res := Simpson(f, 0.0, 10.0, 2.0)
		expected := 20833.33333333333333333333333333

		if res != expected {
			t.Errorf("Result %f does not match expected %f", res, expected)
		}
	}

	{
		res := int(Simpson(f, 0.0, 10.0, 100000.0))
		expected := 19999

		if res != expected {
			t.Errorf("Result %d does not match expected %d", res, expected)
		}
	}
}

func BenchmarkSimpson1(b *testing.B) {

	f := func(x float64) float64 {
		return x * x
	}

	for i := 0; i < b.N; i++ {
		Simpson(f, 0.0, 10.0, 1.0)
	}
}

func BenchmarkSimpson2(b *testing.B) {

	f := func(x float64) float64 {
		return x * x
	}

	for i := 0; i < b.N; i++ {
		Simpson(f, 0.0, 10.0, 2.0)
	}
}

func BenchmarkSimpson200(b *testing.B) {

	f := func(x float64) float64 {
		return x * x
	}

	for i := 0; i < b.N; i++ {
		Simpson(f, 0.0, 10.0, 200.0)
	}
}

func BenchmarkSimpson20000(b *testing.B) {

	f := func(x float64) float64 {
		return x * x
	}

	for i := 0; i < b.N; i++ {
		Simpson(f, 0.0, 10.0, 20000.0)
	}
}
