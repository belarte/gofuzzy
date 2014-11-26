package engine

import "testing"

type f1 struct{}
type f2 struct{}
type f3 struct{}
type f4 struct{}

func (self f1) Compute(_ Object) float64 {
	return 0.25
}

func (self f2) Compute(_ Object) float64 {
	return 0.5
}

func (self f3) Compute(_ Object) float64 {
	return 0.75
}

func (self f4) Compute(_ Object) float64 {
	return 0.125
}

var ruleName string = "(not f1 or f2) and not (f3 and f4)"

func initParserTest() {
	Init()

	knowledgeBase.AddFunction("f1", f1{})
	knowledgeBase.AddFunction("f2", f2{})
	knowledgeBase.AddFunction("f3", f3{})
	knowledgeBase.AddFunction("f4", f4{})

	engine.Fuzzify(Object{})
}

func TestParse(t *testing.T) {
	initParserTest()

	rule, err := Parse(ruleName)

	if err != nil {
		t.Errorf("Error while parsing: ", err)
	}

	expected := 0.75
	result := rule()

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
