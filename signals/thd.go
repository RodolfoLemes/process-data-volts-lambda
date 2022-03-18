package signals

import "math"

func THD(values []float64) float64 {
	sum := 0.0

	for _, value := range values {
		sum += value
	}

	maxValue, _ := max(values)

	harmonicsSum := sum - maxValue

	return 100 * math.Pow(harmonicsSum, 0.5) / (maxValue / 2)
}
