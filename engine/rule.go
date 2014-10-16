package engine

import (
	"math"
)

type Expression func(Object) float64

func BinaryExpressionBuilder(left, right Expression, op Operator) Expression {
	return func(o Object) float64 {
		return op(left(o), right(o))
	}
}

func NegationExpressionBuilder(left Expression) Expression {
	return func(o Object) float64 {
		return 1 - left(o)
	}
}

func ValueExpressionBuilder(f Function) Expression {
	return func(o Object) float64 {
		return f(o)
	}
}

type Operator func(float64, float64) float64

func NewOperator(s string) Operator {
	switch s {
	case "min":
		return minAnd
	case "product":
		return productAnd
	case "max":
		return maxOr
	case "sum":
		return sumOr
	}

	return nil
}

func minAnd(left, right float64) float64 {
	return math.Min(left, right)
}

func productAnd(left, right float64) float64 {
	return left * right
}

func maxOr(left, right float64) float64 {
	return math.Max(left, right)
}

func sumOr(left, right float64) float64 {
	return left + right - left*right
}
