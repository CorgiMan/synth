package synth

const SAMPLERATE = 44100
const dt = 1 / float32(SAMPLERATE)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
