package function

import (
	"errors"
	"math"
)

type Object interface {
	Get(attribute string) float64
}

type MembershipFunction interface {
	Compute(o Object) float64
}

type TrueFunction struct {}

type FalseFunction struct {}

type Function struct {
	attribute string
	set Set
}

func (self TrueFunction) Compute(_ Object) float64 {
	return 1.
}

func (self FalseFunction) Compute(_ Object) float64 {
	return 0.
}

func (self Function) Compute(o Object) float64 {
	value := o.Get(self.attribute)
	return self.set(value)
}

type Set func(float64) float64

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
		case x >= lowSupport && x < lowCore :
			result = (x-lowSupport)/(lowCore-lowSupport);
		case x > highCore && x <= highSupport :
			result = (highSupport-x)/(highSupport-highCore);
		case x >= lowCore && x <= highCore :
			result = 1.0;
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
		case x >= lowSupport && x < lowCore :
			result = (math.Cos((x-lowCore)*math.Pi/(lowCore-lowSupport))+1)/2;
		case x > highCore && x <= highSupport :
			result = (math.Cos((x-highCore)*math.Pi/(highSupport-highCore))+1)/2;
		case x >= lowCore && x <= highCore :
			result = 1.0;
		}

		return result
	}, nil
}
