package engine

import (
	"testing"
)

var leftValues = [...]float64 {
	0.0, 0.25, 0.5, 0.75, 1.0,
}

var rightValues = [...]float64 {
	0.0, 0.5, 1.0,
}

func TestMinAnd (t *testing.T) {
	var expected = [][]float64 {
		{0.0, 0.0, 0.0, 0.0, 0.0,},
		{0.0, 0.25, 0.5, 0.5, 0.5,},
		{0.0, 0.25, 0.5, 0.75, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := MinAnd(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestProductAnd (t *testing.T) {
	var expected = [][]float64 {
		{0.0, 0.0, 0.0, 0.0, 0.0,},
		{0.0, 0.125, 0.25, 0.375, 0.5,},
		{0.0, 0.25, 0.5, 0.75, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := ProductAnd(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestMaxOr (t *testing.T)  {
	var expected = [][]float64 {
		{0.0, 0.25, 0.5, 0.75, 1.0},
		{0.5, 0.5, 0.5, 0.75, 1.0,},
		{1.0, 1.0, 1.0, 1.0, 1.0,},

	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := MaxOr(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestSumOr (t *testing.T)  {
	var expected = [][]float64 {
		{0.0, 0.25, 0.5, 0.75, 1.0},
		{0.5, 0.625, 0.75, 0.875, 1.0,},
		{1.0, 1.0, 1.0, 1.0, 1.0,},

	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := SumOr(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}
