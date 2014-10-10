package synth

import (
	"math"
	"math/rand"
)

type SinGen struct {
	freq  float64
	phase float64
}

func (g *SinGen) Process(out [][]float32) {
	step := g.freq / SAMPLERATE
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phase))
		out[1][i] = out[0][i]
		g.phase += step
		for g.phase > 1 {
			g.phase -= 1
		}
	}

}

type TriGen struct {
	freq  float64
	phase float64
}

func (g *TriGen) Process(out [][]float32) {
	step := g.freq / SAMPLERATE
	for i := range out[0] {
		if g.phase < 0.25 {
			out[0][i] = float32(g.phase) * 4
		} else if g.phase < 0.75 {
			out[0][i] = (0.5 - float32(g.phase)) * 4
		} else {
			out[0][i] = (float32(g.phase) - 1) * 4
		}
		out[1][i] = out[0][i]
		g.phase += step
		for g.phase > 1 {
			g.phase -= 1
		}
	}
}

type SquGen struct {
	freq  float64
	phase float64
}

func (g *SquGen) Process(out [][]float32) {
	step := g.freq / SAMPLERATE
	for i := range out[0] {
		if g.phase < 0.5 {
			out[0][i] = 0.9
		} else {
			out[0][i] = -0.9
		}
		out[1][i] = out[0][i]
		g.phase += step
		for g.phase > 1 {
			g.phase -= 1
		}
	}
}

type SawGen struct {
	freq  float64
	phase float64
}

func (g *SawGen) Process(out [][]float32) {
	step := g.freq / SAMPLERATE
	for i := range out[0] {
		out[0][i] = float32(g.phase)*2 - 1
		out[1][i] = out[0][i]
		g.phase += step
		for g.phase > 1 {
			g.phase -= 1
		}
	}
}

type NoiseGen struct{}

func (g *NoiseGen) Process(out [][]float32) {
	for i := range out[0] {
		out[0][i] = rand.Float32()*2 - 1
		out[1][i] = out[0][i]
	}
}
