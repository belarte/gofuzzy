package engine

import (
	"math"
)

func COGDefuzzifier(min, max float64) float64 {
	sum := 0.0
	weightedSum := 0.0
	step := (max - min) / float64(steps)

	x := min
	ya := computeValue(x)

	for x <= max {
		x += step
		yb := computeValue(x)

		temp := step * (ya + yb) / 2.
		sum += temp
		weightedSum += (x - step*0.5) * temp
		ya = yb
	}

	return weightedSum / sum
}

func MMDefuzzifier(min, max float64) float64 {
	step := (max - min) / float64(steps)

	x := min
	y := computeValue(x)
	maxvalue := y
	start := x
	stop := x

	for x <= max {
		x += step
		y = computeValue(x)

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

func computeValue(x float64) float64 {
	var result float64 = 0.0

	for key, output := range engine.rulesOutputExpression {
		y := math.Min(output.ComputeWithValue(x), engine.rulesOutputValue[key])
		result = math.Max(result, y)
	}

	return result
}
