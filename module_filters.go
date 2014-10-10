package synth

import (
	"math"
)

type HO struct {
	x, v float32
	Freq float32
}

func (ho *HO) Process(out [][]float32) {
	for i := range out[0] {

		ho.v += -dt*(2*math.Pi)*(2*math.Pi)*ho.Freq*ho.Freq*ho.x - dt*out[0][i]
		ho.v *= 0.90
		ho.x += ho.v * dt
		out[0][i] = ho.x * 20000000
		out[1][i] = ho.x * 20000000
	}
}

type LowPass struct {
	Freq       float32
	memL, memR float32
}

func (f *LowPass) Process(out [][]float32) {
	RC := 1 / (f.Freq * 2 * float32(math.Pi))
	a := dt / (RC + dt)
	a = 0.01
	for i := range out[0] {
		out[0][i] = a*out[0][i] + (1-a)*f.memL
		f.memL = out[0][i]
		out[1][i] = a*out[1][i] + (1-a)*f.memR
		f.memR = out[1][i]
	}
}
