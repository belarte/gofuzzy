package utilities

import (
	"math"
)

const (
	Epsilon = 0.000001
)

func CompareEpsilon(a, b float64) bool {
	va := math.Abs(a)
	vb := math.Abs(b)
	diff := math.Abs(va - vb)
	return diff < Epsilon
}

type MinAggregatorF64 struct {
	value float64
}

func NewMinAggregatorF64(val float64) MinAggregatorF64 {
	var result MinAggregatorF64
	result.value = val
	return result
}

func (self *MinAggregatorF64) Push(val float64) {
	if val < self.value {
		self.value = val
	}
}

func (self MinAggregatorF64) Pop() float64 {
	return self.value
}

type MaxAggregatorF64 struct {
	value float64
}

func NewMaxAggregatorF64(val float64) MaxAggregatorF64 {
	var result MaxAggregatorF64
	result.value = val
	return result
}

func (self *MaxAggregatorF64) Push(val float64) {
	if val > self.value {
		self.value = val
	}
}

func (self MaxAggregatorF64) Pop() float64 {
	return self.value
}
