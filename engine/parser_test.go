package engine

import "testing"

var f1 Function = func(_ Object) float64 {
	return 0.25
}

var f2 Function = func(_ Object) float64 {
	return 0.5
}

var f3 Function = func(_ Object) float64 {
	return 0.75
}

var f4 Function = func(_ Object) float64 {
	return 0.125
}

var ruleName string = "(not f1 or f2) and not (f3 and f4)"

func initParserTest() {
	Init()

	knowledgeBase.AddFunction("f1", f1)
	knowledgeBase.AddFunction("f2", f2)
	knowledgeBase.AddFunction("f3", f3)
	knowledgeBase.AddFunction("f4", f4)
}

func TestParse(t *testing.T) {
	initParserTest()

	rule, err := Parse(ruleName)

	if err != nil {
		t.Errorf("Error while parsing: ", err)
	}

	expected := 0.75
	result := rule(Object{})

	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
