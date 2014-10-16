package engine

import (
	"testing"
)

var leftValues = [...]float64{
	0.0, 0.25, 0.5, 0.75, 1.0,
}

var rightValues = [...]float64{
	0.0, 0.5, 1.0,
}

type fakeLeft struct {}
func (self fakeLeft) Compute (_ Object) float64 {
	return 0.25
}

type fakeRight struct {}
func (self fakeRight) Compute (_ Object) float64 {
	return 0.75
}

func TestBinaryExpressionBuilder(t *testing.T) {
	expected := 0.25
	var op Operator = minAnd

	returned := BinaryExpressionBuilder(ValueExpressionBuilder(fakeLeft{}),
		ValueExpressionBuilder(fakeRight{}), op)
	result := returned(Object{})

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestNegationExpressionBuilder(t *testing.T) {
	expected := 0.75

	returned := NegationExpressionBuilder(ValueExpressionBuilder(fakeLeft{}))
	result := returned(Object{})

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestValueExpressionBuilder(t *testing.T) {
	expected := 0.75

	returned := ValueExpressionBuilder(fakeRight{})
	result := returned(Object{})

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestMinAnd(t *testing.T) {
	var expected = [][]float64{
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{0.0, 0.25, 0.5, 0.5, 0.5},
		{0.0, 0.25, 0.5, 0.75, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := minAnd(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestProductAnd(t *testing.T) {
	var expected = [][]float64{
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{0.0, 0.125, 0.25, 0.375, 0.5},
		{0.0, 0.25, 0.5, 0.75, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := productAnd(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestMaxOr(t *testing.T) {
	var expected = [][]float64{
		{0.0, 0.25, 0.5, 0.75, 1.0},
		{0.5, 0.5, 0.5, 0.75, 1.0},
		{1.0, 1.0, 1.0, 1.0, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := maxOr(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}

func TestSumOr(t *testing.T) {
	var expected = [][]float64{
		{0.0, 0.25, 0.5, 0.75, 1.0},
		{0.5, 0.625, 0.75, 0.875, 1.0},
		{1.0, 1.0, 1.0, 1.0, 1.0},
	}

	for i, lVal := range leftValues {
		for j, rVal := range rightValues {
			result := sumOr(lVal, rVal)
			exp := expected[j][i]

			if result != exp {
				t.Errorf("Expected: %v, got: %v", exp, result)
			}
		}
	}
}
