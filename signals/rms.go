package signals

func CalculateFP(
	tensionSignal *Signal,
	currentSignal *Signal,
	size int,
) (FP float64) {
	S := tensionSignal.GetRMS() * currentSignal.GetRMS()

	P := 0.0
	for i := range tensionSignal.GetValues()[0 : size-1] {
		P = tensionSignal.GetValues()[i] * currentSignal.GetValues()[i]
	}

	FP = P / S / float64(size)

	return
}
