package engine

const (
	NB_STEPS = 1000
)

type Defuzzifier interface {
	Compute() float64
}

type COGDefuzzifier struct {
	min, max float64
}

func (self COGDefuzzifier) Compute() float64 {
	result := 0.0
	step := (self.max - self.min) / NB_STEPS

	x := self.min
	ya := engine.computeValue(x)

	for x <= self.max {
		x += step
		yb := engine.computeValue(x)

		result += step * (ya + yb) / 2
		ya = yb
	}

	return result
}

type MMDefuzzifier struct {
	min, max float64
}

func (self MMDefuzzifier) Compute() float64 {
	step := (self.max - self.min) / NB_STEPS

	x := self.min
	y := engine.computeValue(x)
	maxvalue := y
	start := x
	stop := x

	for x <= self.max {
		x += step
		y = engine.computeValue(x)

		if y > maxvalue {
			maxvalue = y
			start = x
			stop = x
		} else if y == maxvalue {
			stop = x
		}
	}

	return (stop + start) / 2
}
