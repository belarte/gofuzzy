package engine

import (
	"math"
)

func BinaryExpressionBuilder(left, right Function, op Operator) Function {
	return func(o Object) float64 {
		return op(left(o), right(o))
	}
}

func NegationExpressionBuilder(left Function) Function {
	return func(o Object) float64 {
		return 1 - left(o)
	}
}

func ValueExpressionBuilder(f Function) Function {
	return func(o Object) float64 {
		return f(o)
	}
}

type Operator func(float64, float64) float64

func MinAnd(left, right float64) float64 {
	return math.Min(left, right)
}

func ProductAnd(left, right float64) float64 {
	return left * right
}

func MaxOr(left, right float64) float64 {
	return math.Max(left, right)
}

func SumOr(left, right float64) float64 {
	return left + right - left*right
}
