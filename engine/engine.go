package engine

import (
	"math"
)

type Engine struct {
	functionsOutput       map[string]float64
	rulesOutputValue      map[string]float64
	rulesOutputExpression map[string]MembershipFunction
	result                float64
	isComputed            bool
}

func NewEngine() Engine {
	var result Engine
	result.functionsOutput = make(map[string]float64)
	result.rulesOutputValue = make(map[string]float64)
	result.rulesOutputExpression = make(map[string]MembershipFunction)
	return result
}

func (self Engine) AddFunction(name string, value float64) {
	self.functionsOutput[name] = value
}

func (self Engine) FunctionOutput(name string) float64 {
	return self.functionsOutput[name]
}

func (self Engine) Compute() {
	self.Fuzzify(Object{})
	self.Infer()
	self.Defuzzify()
}

func (self Engine) Fuzzify(obj Object) {
	for key := range self.functionsOutput {
		delete(self.functionsOutput, key)
	}

	for key, function := range knowledgeBase.Functions() {
		self.AddFunction(key, function.Compute(obj))
	}
}

func (self Engine) Infer() {
	for key := range self.rulesOutputValue {
		delete(self.rulesOutputValue, key)
		delete(self.rulesOutputExpression, key)
	}

	for key, rule := range knowledgeBase.Rules() {
		value, expr := rule.Compute()
		self.rulesOutputValue[key] = value
		self.rulesOutputExpression[key] = expr
	}
}

func (self Engine) Defuzzify() {
	min, max := self.getMinMax()

	self.result = Operators[defuzzyOperator](min, max)

	self.isComputed = true
}

func (self Engine) getMinMax() (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, function := range self.rulesOutputExpression {
		temp := function.Min()
		if temp < min {
			min = temp
		}

		temp = function.Max()
		if temp > max {
			max = temp
		}
	}

	return min, max
}
