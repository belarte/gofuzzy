package engine

import (
	"testing"
	"github.com/belarte/gofuzzy/utilities"
)

const (
	min        = 0.
	minSupport = 10.
	minCore    = 20.
	maxCore    = 30.
	maxSupport = 40.
	max        = 50.
)

var values = [...]float64{
	min,
	minSupport - utilities.Epsilon,
	minSupport,
	minSupport + utilities.Epsilon,
	minSupport + (minCore-minSupport)/4,
	minSupport + (minCore-minSupport)/2,
	minSupport + 3*(minCore-minSupport)/4,
	minCore - utilities.Epsilon,
	minCore,
	minCore + utilities.Epsilon,
	(minCore + maxCore) / 2,
	maxCore - utilities.Epsilon,
	maxCore,
	maxCore + utilities.Epsilon,
	maxCore + (maxSupport-maxCore)/4,
	maxCore + (maxSupport-maxCore)/2,
	maxCore + 3*(maxSupport-maxCore)/4,
	maxSupport - utilities.Epsilon,
	maxSupport,
	maxSupport + utilities.Epsilon,
	max,
}

func TestBooleanSetBuilder(t *testing.T) {
	var expected = []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	function, err := BooleanSetBuilder(minCore, maxCore)

	if err != nil {
		t.Error("Error while building the function.")
	}

	for i, value := range values {
		result := function(value)
		if !utilities.CompareEpsilon(result, expected[i]) {
			t.Errorf("Expected: %v, got: %v", expected[i], result)
		}
	}
}

func TestTrapezoidalSetBuilder(t *testing.T) {
	var expected = []float64{
		0, 0, 0, 0, 0.25, 0.5, 0.75, 1, 1, 1, 1, 1, 1, 1, 0.75, 0.5, 0.25, 0, 0, 0, 0,
	}

	function, err := TrapezoidalSetBuilder(minSupport, minCore, maxCore, maxSupport)

	if err != nil {
		t.Error("Error while building the function.")
	}

	for i, value := range values {
		result := function(value)
		if !utilities.CompareEpsilon(result, expected[i]) {
			t.Errorf("Expected: %v, got: %v", expected[i], result)
		}
	}
}

func TestSinusoidalSetBuilder(t *testing.T) {
	var expected = []float64{
		0, 0, 0, 0, 0.146447, 0.5, 0.853553, 1, 1, 1, 1, 1, 1, 1, 0.853553, 0.5, 0.146447, 0, 0, 0, 0,
	}

	function, err := SinusoidalSetBuilder(minSupport, minCore, maxCore, maxSupport)

	if err != nil {
		t.Error("Error while building the function.")
	}

	for i, value := range values {
		result := function(value)
		if !utilities.CompareEpsilon(result, expected[i]) {
			t.Errorf("Expected: %v, got: %v", expected[i], result)
		}
	}
}
