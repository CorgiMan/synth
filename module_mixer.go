package synth

import "fmt"

type Vol struct {
	Fact float32
}

func (m *Vol) Process(out [][]float32) {
	for i := range out[0] {
		out[0][i] *= m.Fact
		out[1][i] *= m.Fact
	}
}

type Mixer struct {
	inputs []Module
	faders []float32
	buffer [][]float32
}

func NewMixer() *Mixer {
	return &Mixer{buffer: make([][]float32, 2)}
}

func (m *Mixer) AddInput(s Module, faderval float32) {
	m.inputs = append(m.inputs, s)
	m.faders = append(m.faders, faderval)
}

func (m *Mixer) Process(out [][]float32) {
	if len(out[0]) > len(m.buffer[0]) {
		m.buffer[0] = append(m.buffer[0], out[0][len(m.buffer[0]):]...)
		m.buffer[1] = append(m.buffer[1], out[1][len(m.buffer[1]):]...)
	}

	for x := range out[0] {
		out[0][x] = 0
		out[1][x] = 0
	}

	for i, in := range m.inputs {
		in.Process(m.buffer)
		for x := range out[0] {
			out[0][x] += m.buffer[0][x] * m.faders[i]
			out[1][x] += m.buffer[1][x] * m.faders[i]
		}
	}

	fmt.Println(out[0][100])
}
