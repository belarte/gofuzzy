package engine

import (
	"log"
	"math"
)

type Rule struct {
	name   string
	input  Expression
	output MembershipFunction
}

func (self Rule) Compute() (float64, MembershipFunction) {
	return self.input(), self.output
}

type Expression func() float64

func AndExpressionBuilder(left, right Expression) Expression {
	return func() float64 {
		op := Operators[andOperator]
		return op(left(), right())
	}
}

func OrExpressionBuilder(left, right Expression) Expression {
	return func() float64 {
		op := Operators[orOperator]
		return op(left(), right())
	}
}

func NegationExpressionBuilder(left Expression) Expression {
	return func() float64 {
		return 1 - left()
	}
}

func ValueExpressionBuilder(name string) Expression {
	return func() float64 {
		val, check := engine.FunctionOutput(name)
		if !check {
			log.Println("Function `", name, "' not found.")
		}
		return val
	}
}

type Operator func(float64, float64) float64

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
