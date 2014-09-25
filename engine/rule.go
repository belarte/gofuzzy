package engine

import (
	"math"
)

type Expression func(o Object) float64

func BinaryExpressionBuilder (left, right Expression, op Operator) Expression {
	return func (o Object) float64 {
		return op(left(o), right(o))
	}
}

func NegationExpressionBuilder (left Expression) Expression {
	return func (o Object) float64 {
		return 1 - left(o)
	}
}

func ValueExpressionBuilder (f Function) Expression {
	return func (o Object) float64 {
		return f.Compute(o)
	}
}

type Operator func (left, right float64) float64

func MinAnd (left, right float64) float64 {
	return math.Min (left,right)
}

func ProductAnd (left, right float64) float64 {
	return left * right
}

func MaxOr (left, right float64) float64 {
	return math.Max (left,right)
}

func SumOr (left, right float64) float64 {
	return left + right - left * right
}
