package engine

import (
	"github.com/belarte/gofuzzy/utilities"
	"testing"
)

var (
	testObject Object = NewObject(map[string]float64{"speed": 105, "distance": 45})
)

func SetUp() {
	Init()

	if fun, err := NewMembershipFunction("speed", "trapezoidal", []float64{0, 0, 70, 90}); err == nil {
		knowledgeBase.AddFunction("slow", fun)
	}
	if fun, err := NewMembershipFunction("speed", "trapezoidal", []float64{70, 90, 90, 110}); err == nil {
		knowledgeBase.AddFunction("good", fun)
	}
	if fun, err := NewMembershipFunction("speed", "trapezoidal", []float64{90, 110, 130, 130}); err == nil {
		knowledgeBase.AddFunction("fast", fun)
	}
	if fun, err := NewMembershipFunction("distance", "trapezoidal", []float64{0, 0, 30, 50}); err == nil {
		knowledgeBase.AddFunction("close", fun)
	}
	if fun, err := NewMembershipFunction("distance", "trapezoidal", []float64{30, 50, 100, 100}); err == nil {
		knowledgeBase.AddFunction("far", fun)
	}
	if fun, err := NewMembershipFunction("response", "trapezoidal", []float64{-90, -90, -90, 0}); err == nil {
		knowledgeBase.AddFunction("decelerate", fun)
		if expr, err := Parse("fast or close"); err == nil {
			knowledgeBase.AddRule("decelerate", Rule{"decelerate", expr, fun})
		}
	}
	if fun, err := NewMembershipFunction("response", "trapezoidal", []float64{-90, 0, 0, 90}); err == nil {
		knowledgeBase.AddFunction("keep", fun)
		if expr, err := Parse("good and far"); err == nil {
			knowledgeBase.AddRule("keep", Rule{"keep", expr, fun})
		}
	}
	if fun, err := NewMembershipFunction("response", "trapezoidal", []float64{0, 90, 90, 90}); err == nil {
		knowledgeBase.AddFunction("accelerate", fun)
		if expr, err := Parse("slow and far"); err == nil {
			knowledgeBase.AddRule("accelerate", Rule{"accelerate", expr, fun})
		}
	}
}

func TestCompute(t *testing.T) {
	SetUp()
	t.Error("TODO")
}

func TestFuzzify(t *testing.T) {
	engine.Fuzzify(testObject)

	functionsName := [...]string{"slow", "good", "fast", "close", "far"}
	expectedOutput := [...]float64{0, 0.25, 0.75, 0.25, 0.75}

	for i, name := range functionsName {
		output, check := engine.FunctionOutput(name)
		expected := expectedOutput[i]
		if !check {
			t.Error("Error accessing function `", name, "'")
		}
		if output != expected {
			t.Errorf("Expected: %v, got: %v", expected, output)
		}
	}
}

func TestInfer(t *testing.T) {
	engine.Fuzzify(testObject)
	engine.Infer()

	rulesName := [...]string{"keep", "accelerate", "decelerate"}
	expectedValues := [...]float64{0.25, 0, 0.75}

	for i, rule := range rulesName {
		output, check := engine.rulesOutputValue[rule]
		if !check {
			t.Error("Error accessing rule `", rule, "'")
		}

		expected := expectedValues[i]
		if output != expected {
			t.Errorf("Expected: %v, got: %v", expected, output)
		}
	}
}

func TestDefuzzify(t *testing.T) {
	engine.Fuzzify(testObject)
	engine.Infer()

	defuzzyOperator = "cog"
	engine.Defuzzify()
	if !engine.isComputed {
		t.Error("Return value not computed...")
	}

	output := engine.result
	expected := -25.3125
	if !utilities.CompareEpsilon(output, expected) {
		t.Errorf("Expected: %v, got: %v", expected, output)
	}

	defuzzyOperator = "mm"
	engine.Defuzzify()
	if !engine.isComputed {
		t.Error("Return value not computed...")
	}

	output = engine.result
	expected = -78.75
	if !utilities.CompareEpsilon(output, expected) {
		t.Errorf("Expected: %v, got: %v", expected, output)
	}
}
