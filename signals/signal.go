package signals

import (
	"math"

	"github.com/mjibson/go-dsp/fft"
)

type Signal struct {
	id              string
	values          []float64
	frequencySample float64
	Frequency       float64
	FftProperties   *fftProperties
}

func New(id string) *Signal {
	return &Signal{
		id:              id,
		values:          make([]float64, 0),
		frequencySample: 3333.3333,
		FftProperties:   nil,
	}
}

func (signal *Signal) AddValues(values ...float64) {
	signal.values = append(values, values...)
	Ploter(values, values, "values.png")
}

func (signal *Signal) GetValues() []float64 {
	return signal.values
}

func (signal *Signal) GetRMS() float64 {
	sum := 0.0
	for _, value := range signal.values {
		sum += math.Pow(value, 2)
	}
	return math.Sqrt(sum / float64(len(signal.values)))
}

func (signal *Signal) CalculateFftProperties() *fftProperties {
	fftValues := fft.FFTReal(signal.values)
	floatLen := float64(len(signal.values))

	halfLenFftValues := len(signal.values)/2 + 1

	filteredFftValues := make([]complex128, halfLenFftValues)
	absValues := make([]float64, halfLenFftValues)
	filteredFftValues = fftValues[0:halfLenFftValues]
	fStep := signal.frequencySample / floatLen
	frequencies := linspace(0, (floatLen-1)*fStep, len(signal.values))

	// normalizing and abs the data
	var realValue float64
	var imagValue float64
	for i := range filteredFftValues {
		realValue = real(filteredFftValues[i])
		imagValue = imag(filteredFftValues[i])
		absValues[i] = 2 * abs(realValue, imagValue) / floatLen
	}

	Ploter(absValues, frequencies[0:100], "abs.png")
	_, index := max(absValues)
	trueFrequency := frequencies[index]
	signal.Frequency = trueFrequency

	thd := THD(absValues)

	properties := &fftProperties{
		Values:      filteredFftValues,
		AbsValues:   absValues,
		Frequencies: frequencies,
		THD:         thd,
	}

	signal.FftProperties = properties

	return properties
}
