package synth

import (
	"bufio"
	bin "encoding/binary"
	"fmt"
	"os"
)

type FilePlayer struct {
	data [][]float32
}

func NewFilePlayer(file string) *FilePlayer {
	p := &FilePlayer{}
	wav := ReadWavData(file)
	p.data = channeldata(wav.data)
	return p
}

func (p *FilePlayer) Process(out [][]float32) {
	for i := range out[0] {
		if i > len(p.data[0]) {
			out[0][i] = 0
			out[1][i] = 0
		} else {
			out[0][i] = p.data[0][i]
			out[1][i] = p.data[1][i]
		}
	}

	if len(out[0]) < len(p.data[0]) {
		p.data[0] = p.data[0][len(out[0]):]
		p.data[1] = p.data[1][len(out[1]):]
	} else {
		p.data[0] = p.data[0][0:0]
		p.data[1] = p.data[1][0:0]
	}
}

type WavData struct {
	bChunkID  [4]byte // B
	ChunkSize uint32  // L
	bFormat   [4]byte // B

	bSubchunk1ID  [4]byte // B
	Subchunk1Size uint32  // L
	AudioFormat   uint16  // L
	NumChannels   uint16  // L
	SampleRate    uint32  // L
	ByteRate      uint32  // L
	BlockAlign    uint16  // L
	BitsPerSample uint16  // L

	bSubchunk2ID  [4]byte // B
	Subchunk2Size uint32  // L
	data          []byte  // L
}

func ReadWavData(fn string) (wav WavData) {
	ftotal, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Error opening\n")
	}
	file := bufio.NewReader(ftotal)

	bin.Read(file, bin.BigEndian, &wav.bChunkID)
	bin.Read(file, bin.LittleEndian, &wav.ChunkSize)
	bin.Read(file, bin.BigEndian, &wav.bFormat)

	bin.Read(file, bin.BigEndian, &wav.bSubchunk1ID)
	bin.Read(file, bin.LittleEndian, &wav.Subchunk1Size)
	bin.Read(file, bin.LittleEndian, &wav.AudioFormat)
	bin.Read(file, bin.LittleEndian, &wav.NumChannels)
	bin.Read(file, bin.LittleEndian, &wav.SampleRate)
	bin.Read(file, bin.LittleEndian, &wav.ByteRate)
	bin.Read(file, bin.LittleEndian, &wav.BlockAlign)
	bin.Read(file, bin.LittleEndian, &wav.BitsPerSample)

	bin.Read(file, bin.BigEndian, &wav.bSubchunk2ID)
	bin.Read(file, bin.LittleEndian, &wav.Subchunk2Size)

	wav.data = make([]byte, wav.Subchunk2Size)
	bin.Read(file, bin.LittleEndian, &wav.data)

	// fmt.Printf("\n")
	// fmt.Printf("ChunkID*: %s\n", wav.bChunkID)
	// fmt.Printf("ChunkSize: %d\n", wav.ChunkSize)
	// fmt.Printf("Format: %s\n", wav.bFormat)
	// fmt.Printf("\n")
	// fmt.Printf("Subchunk1ID: %s\n", wav.bSubchunk1ID)
	// fmt.Printf("Subchunk1Size: %d\n", wav.Subchunk1Size)
	// fmt.Printf("AudioFormat: %d\n", wav.AudioFormat)
	// fmt.Printf("NumChannels: %d\n", wav.NumChannels)
	// fmt.Printf("SampleRate: %d\n", wav.SampleRate)
	// fmt.Printf("ByteRate: %d\n", wav.ByteRate)
	// fmt.Printf("BlockAlign: %d\n", wav.BlockAlign)
	// fmt.Printf("BitsPerSample: %d\n", wav.BitsPerSample)
	// fmt.Printf("\n")
	// fmt.Printf("Subchunk2ID: %s\n", wav.bSubchunk2ID)
	// fmt.Printf("Subchunk2Size: %d\n", wav.Subchunk2Size)
	// fmt.Printf("NumSamples: %d\n", wav.Subchunk2Size/uint32(wav.NumChannels)/uint32(wav.BitsPerSample/8))
	// fmt.Printf("\ndata: %v\n", len(wav.data))
	// fmt.Printf("\n\n")

	return
}

func btou(b []byte) (u []uint16) {
	u = make([]uint16, len(b)/2)
	for i, _ := range u {
		val := uint16(b[i*2])
		val += uint16(b[i*2+1]) << 8
		u[i] = val
	}
	return
}

func btoi16(b []byte) (u []int16) {
	u = make([]int16, len(b)/2)
	for i, _ := range u {
		val := int16(b[i*2])
		val += int16(b[i*2+1]) << 8
		u[i] = val
	}
	return
}

func btof32(b []byte) (f []float32) {
	u := btoi16(b)
	f = make([]float32, len(u))
	for i, v := range u {
		f[i] = float32(v) / float32(32768)
	}
	return
}

func utob(u []uint16) (b []byte) {
	b = make([]byte, len(u)*2)
	for i, val := range u {
		lo := byte(val)
		hi := byte(val >> 8)
		b[i*2] = lo
		b[i*2+1] = hi
	}
	return
}

func channeldata(b []byte) (f [][]float32) {
	u := btoi16(b)
	f = make([][]float32, 2)
	f[0] = make([]float32, len(u)/2)
	f[1] = make([]float32, len(u)/2)
	for i, v := range u {
		if i%2 == 0 {
			f[0][i/2] = float32(v) / float32(32768)
		} else {
			f[1][i/2] = float32(v) / float32(32768)
		}
	}
	return
}
