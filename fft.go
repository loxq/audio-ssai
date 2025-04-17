package main

import (
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

func containsBeep(samples []float64) bool {
	spectrum := fft.FFTReal(samples)
	freqResolution := float64(SampleRate) / float64(len(samples))
	targetBin := int(FrequencyToDetect / freqResolution)

	magnitude := cmplx.Abs(spectrum[targetBin])
	total := 0.0
	for _, c := range spectrum {
		total += cmplx.Abs(c)
	}

	return magnitude/total > Threshold
}
