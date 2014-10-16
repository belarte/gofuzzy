package engine

import (
	"errors"
	"math"
)

type Object struct {
	attributes map[string]float64
}

func (self Object) Get(name string) (float64, bool) {
	res, check := self.attributes[name]
	return res, check
}

type Function interface {
	Compute (o Object) float64
}

type MembershipFunction struct {
	attribute string
	shape string
	params []float64
	set Set
}

type TrueFunction struct {}
type FalseFunction struct {}

func NewMembershipFunction (attribute, shape string, params []float64) (MembershipFunction, error) {
	set, err := NewSet(shape, params)
	return MembershipFunction {attribute, shape, params, set}, err
}

func (self MembershipFunction) Compute (o Object) float64 {
	value, _ := o.Get(self.attribute)
	return self.set(value)
}

func (self TrueFunction) Compute (o Object) float64 {
	return 1.0
}

func (self FalseFunction) Compute (o Object) float64 {
	return 0.0
}

type Set func(float64) float64

func NewSet(op string, params []float64) (Set, error) {
	switch op {
	case "boolean":
		if len(params) != 2 {
			return nil, errors.New("Wrong number of arguments")
		}
		return BooleanSetBuilder(params[0], params[1])
	case "trapezoidal":
		if len(params) != 4 {
			return nil, errors.New("Wrong number of arguments")
		}
		return TrapezoidalSetBuilder(params[0], params[1], params[2], params[3])
	case "sinusoidal":
		if len(params) != 4 {
			return nil, errors.New("Wrong number of arguments")
		}
		return SinusoidalSetBuilder(params[0], params[1], params[2], params[3])
	}
	return nil, errors.New("Unable to generate set")
}

func BooleanSetBuilder(lowCore, highCore float64) (Set, error) {
	if lowCore > highCore {
		return nil, errors.New("Error while constructing BooleanSet.\nIntervals are not properly set.")
	}

	return func(x float64) float64 {
		result := 0.

		if x >= lowCore && x <= highCore {
			result = 1.
		}

		return result
	}, nil
}

func TrapezoidalSetBuilder(lowSupport, lowCore, highCore, highSupport float64) (Set, error) {
	if lowSupport > lowCore || lowCore > highCore || highCore > highSupport {
		return nil, errors.New("Error while constructing TrapezoidalSet.\nIntervals are not properly set.")
	}

	return func(x float64) float64 {
		result := 0.

		switch {
		case x >= lowSupport && x < lowCore:
			result = (x - lowSupport) / (lowCore - lowSupport)
		case x > highCore && x <= highSupport:
			result = (highSupport - x) / (highSupport - highCore)
		case x >= lowCore && x <= highCore:
			result = 1.0
		}

		return result
	}, nil
}

func SinusoidalSetBuilder(lowSupport, lowCore, highCore, highSupport float64) (Set, error) {
	if lowSupport > lowCore || lowCore > highCore || highCore > highSupport {
		return nil, errors.New("Error while constructing SinusoidalSet.\nIntervals are not properly set.")
	}

	return func(x float64) float64 {
		result := 0.

		switch {
		case x >= lowSupport && x < lowCore:
			result = (math.Cos((x-lowCore)*math.Pi/(lowCore-lowSupport)) + 1) / 2
		case x > highCore && x <= highSupport:
			result = (math.Cos((x-highCore)*math.Pi/(highSupport-highCore)) + 1) / 2
		case x >= lowCore && x <= highCore:
			result = 1.0
		}

		return result
	}, nil
}
