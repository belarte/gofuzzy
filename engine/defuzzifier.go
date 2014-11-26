package engine

type Defuzzifier interface {
	Compute() float64
}

type COGDefuzzifier struct{}

func (self COGDefuzzifier) Compute() float64 {
	return 0.
}

type MMDefuzzifier struct{}

func (self MMDefuzzifier) Compute() float64 {
	return 0.
}
