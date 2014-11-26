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

func TestBinaryExpressionBuilder(t *testing.T) {
	expected := 0.25

	engine.AddFunction("fakeLeft", 0.25)
	engine.AddFunction("fakeRight", 0.75)

	returned := AndExpressionBuilder(ValueExpressionBuilder("fakeLeft"),
		ValueExpressionBuilder("fakeRight"))
	result := returned()

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestNegationExpressionBuilder(t *testing.T) {
	expected := 0.75

	returned := NegationExpressionBuilder(ValueExpressionBuilder("fakeLeft"))
	result := returned()

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestValueExpressionBuilder(t *testing.T) {
	expected := 0.75

	returned := ValueExpressionBuilder("fakeRight")
	result := returned()

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
