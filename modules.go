package synth

type Module interface {
	Process(out [][]float32)
}

type MonoModule interface {
	Process(out []float32)
}

// STACK MODULE
type ModuleStack struct {
	modules []Module
}

func (r *ModuleStack) AddModule(m Module) {
	r.modules = append(r.modules, m)
}

func (r *ModuleStack) Process(out [][]float32) {
	for _, m := range r.modules {
		m.Process(out)
	}
}

// STEREO MODULE
type StereoModule struct {
	left, right MonoModule
}

func (m *StereoModule) Process(out [][]float32) {
	m.left.Process(out[0])
	m.right.Process(out[1])
}

type NoiseReductor struct {
	level float32
}

func (m *NoiseReductor) Process(out [][]float32) {
	for i := range out[0] {
		if out[0][i] < m.level {
			out[0][i] = 0
		}
	}
}
