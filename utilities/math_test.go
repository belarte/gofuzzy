package utilities

import (
	"testing"
)

var (
	values 		= [...]float64 { 7., 5., 3., 1., 6., 9., 8., 4., 2. }
	expectedMin = [...]float64 { 7., 5., 3., 1., 1., 1., 1., 1., 1. }
	expectedMax = [...]float64 { 7., 7., 7., 7., 7., 9., 9., 9., 9. }
)

func TestMinAggregatorF64(t *testing.T) {
	aggregator := MinAggregatorF64{100.}

	for i, value := range values {
		aggregator.Push(value)

		result := aggregator.Pop()
		if !CompareEpsilon(result, expectedMin[i]) {
			t.Errorf("Expected: %v, got: %v", expectedMin[i], result)
		}
	}
}

func TestMaxAggregatorF64(t *testing.T) {
	aggregator := MaxAggregatorF64{0.}

	for i, value := range values {
		aggregator.Push(value)

		result := aggregator.Pop()
		if !CompareEpsilon(result, expectedMax[i]) {
			t.Errorf("Expected: %v, got: %v", expectedMax[i], result)
		}
	}
}
