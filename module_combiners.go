package synth

type Sub struct {
	m1, m2 Module
	buffer [][]float32
}

func NewSub(m1, m2 Module) *Sub {
	return &Sub{m1, m2, make([][]float32, 2)}
}

func (m *Sub) Process(out [][]float32) {
	if len(out[0]) > len(m.buffer[0]) {
		m.buffer[0] = append(m.buffer[0], out[0][len(m.buffer[0]):]...)
		m.buffer[1] = append(m.buffer[1], out[1][len(m.buffer[1]):]...)
	}

	m.m1.Process(out)
	m.m2.Process(m.buffer)

	for i := range out[0] {
		out[0][i] -= m.buffer[0][i]
		out[1][i] -= m.buffer[1][i]
	}
}

type Mul struct {
	m1, m2 Module
	buffer [][]float32
}

func NewMul(m1, m2 Module) *Sub {
	return &Sub{m1, m2, make([][]float32, 2)}
}

func (m *Mul) Process(out [][]float32) {
	if len(out[0]) > len(m.buffer[0]) {
		m.buffer[0] = append(m.buffer[0], out[0][len(m.buffer[0]):]...)
		m.buffer[1] = append(m.buffer[1], out[1][len(m.buffer[1]):]...)
	}

	m.m1.Process(out)
	m.m2.Process(m.buffer)

	for i := range out[0] {
		out[0][i] *= m.buffer[0][i]
		out[1][i] *= m.buffer[1][i]
	}
}
