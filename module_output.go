package synth

import (
	"code.google.com/p/portaudio-go/portaudio"
)

type OutputModule struct {
	*portaudio.Stream
	module Module
}

func NewOutputModule(m Module, sampleRate float64) *OutputModule {
	out := &OutputModule{module: m}
	var err error
	out.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, out.Process)
	chk(err)
	return out
}

func (outm *OutputModule) Process(out [][]float32) {
	outm.module.Process(out)
}
