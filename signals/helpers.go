package signals

import "math"

func BuildSin(Vmax float64, frequency float64, phase float64) *Signal {
	sin := func(t float64) float64 {
		return Vmax * math.Sin(frequency*2*math.Pi*t+phase)
	}

	arr := arange2(0, 0.1, 3*math.Pow10(-4))

	points := make([]float64, len(arr))

	for i, element := range arr {
		points[i] = sin(element)
	}

	return &Signal{
		id:     "sin",
		values: points,
	}
}

func arange2(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N)
	for x := range rnge {
		rnge[x] = start + step*float64(x)
	}
	return rnge
}

func linspace(start, stop float64, num int) []float64 {
	step := 0.

	if num == 0 {
		return []float64{}
	}
	step = (stop - start) / float64(num)
	r := make([]float64, num, num)
	for i := 0; i < num; i++ {
		r[i] = start + float64(i)*step
	}
	return r
}

func max(slice []float64) (float64, int) {
	value := 0.0
	index := 0

	for i := range slice {
		if slice[i] > value {
			value = slice[i]
			index = i
		}
	}

	return value, index
}

func abs(real float64, imag float64) float64 {
	real = math.Pow(real, 2)
	imag = math.Pow(imag, 2)
	return math.Sqrt(real + imag)
}
