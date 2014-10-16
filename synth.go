package synth

const SAMPLERATE = 44100
const dt = 1 / float32(SAMPLERATE)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}


type Synth struct {
    *portaudio.Stream
    ModuleGraph
    buffers map[*Module]Buffer
}

type Buffer [][]byte

func NewSynth(sampleRate float64) *Synth {
    s := &Synth{}
    s.Stream, err := portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.Process)
    chk(err)
    return s
}

func (s *Synth) Process(out [][]byte) {
    // set up buffers if they don't already exist
    
}

type E interface{}

type ModuleGraph struct {
    output *Module
    connections map[*Module]Set
}

func NewModuleGraph(vs []*Module, edgs map[int]map[int]bool, out_module *Module) *ModuleGraph {
    g := &ModuleGraph{ouput: out_module}
    
    g.connections = map[*Module]Set
    for i1 := range edgs {
        g.connections[vs[i1]] = Set{}
        for i2 := range edgs[i1] {
            g.connections[vs[i1]].Add(vs[i2])
        }
    }
}

type Set map[E]E

func (s *Set) Add(elt E) {
    is s==nil {
        s = Set{}
    }
    s[elt] = nil
}

func (s *Set) Delete(elt E) /*bool*/ {
    // _, ok := s[elt]
    delete(s, elt)
    // return ok
}

func (s *Set) Has(elt E) bool {
    _, ok := s[elt]
    return ok
}
