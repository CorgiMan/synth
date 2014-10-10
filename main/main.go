package main

import (
	"code.google.com/p/portaudio-go/portaudio"
	"fmt"
	"github.com/CorgiMan/synth"
	"os"
	"os/signal"
	"time"
)

var _ = fmt.Println

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	out := GetStream()
	defer out.Close()
	chk(out.Start())

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-sig:
			chk(out.Stop())
			return
		default:
		}
	}
}

func GetStream() *portaudio.Stream {
	F := float32(2400)
	s1 := &synth.ModuleStack{}
	s1.AddModule(synth.NewFilePlayer("song.wav"))
	s1.AddModule(&synth.HO{Freq: F})

	s2 := &synth.ModuleStack{}
	s2.AddModule(synth.NewFilePlayer("song.wav"))
	s2.AddModule(&synth.HO{Freq: F / 2})
	s2.AddModule(&synth.Vol{Fact: 0.1})

	s4 := &synth.ModuleStack{}
	s4.AddModule(synth.NewFilePlayer("song.wav"))
	s4.AddModule(&synth.HO{Freq: F * 2})

	s3 := &synth.ModuleStack{}
	s3.AddModule(synth.NewSub(synth.NewSub(s1, s2), s4))
	// s3.AddModule(s4)
	s3.AddModule(&synth.Vol{Fact: 10})

	// s3 := &ModuleStack{}
	// s3.AddModule(&NoiseGen{})
	// s3.AddModule(&HO{freq: 330})
	// s3.AddModule(&HO{freq: 550})
	// s3.AddModule(&HO{freq: 880})
	// s3.AddModule(&LowPass{freq: 6000})

	mixer := synth.NewMixer()
	mixer.AddInput(s3, 0.1)
	// mixer.AddInput(s1, 1)

	out := synth.NewOutputModule(mixer, 44100)
	return out.Stream
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
