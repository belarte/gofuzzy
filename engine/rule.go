package engine

import (
	"math"
)

type Rule struct {
	name   string
	input  Expression
	output *MembershipFunction
}

func (self Rule) Compute() (float64, MembershipFunction) {
	return self.input(), *self.output
}

type Expression func() float64

func AndExpressionBuilder(left, right Expression) Expression {
	return func() float64 {
		return engine.andOperator(left(), right())
	}
}

func OrExpressionBuilder(left, right Expression) Expression {
	return func() float64 {
		return engine.orOperator(left(), right())
	}
}

func NegationExpressionBuilder(left Expression) Expression {
	return func() float64 {
		return 1 - left()
	}
}

func ValueExpressionBuilder(name string) Expression {
	return func() float64 {
		return engine.FunctionOutput(name)
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
