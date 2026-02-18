package sound

import (
	"encoding/binary"
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/ebitengine/oto/v3"
)

type Effect struct {
	time      int
	isPlaying bool

	attackSamples  int
	sustainSamples int
	releaseSamples int

	frequency float32
	amplitude float32
	wave      waveform

	player *oto.Player
}

type waveform int

const (
	sinWave waveform = iota
	triangleWave
	squareWave
	sawtoothWave
	noiseWave
)

const maxAmplitude = 0.25

func NewEffect() *Effect {
	return (&Effect{}).Volume(1).Frequency(440).Sustain(time.Millisecond * 250)
}

func (osc *Effect) Volume(amplitude float32) *Effect {
	osc.amplitude = amplitude * maxAmplitude
	return osc
}

func (osc *Effect) Frequency(frequency float32) *Effect {
	osc.frequency = frequency
	return osc
}

func (osc *Effect) Sustain(duration time.Duration) *Effect {
	osc.ASR(time.Millisecond*3, duration, time.Millisecond*3)
	return osc
}

func (osc *Effect) ASR(attack, sustain, release time.Duration) *Effect {
	osc.attackSamples = int(attack.Seconds() * sampleRate)
	osc.sustainSamples = int(sustain.Seconds() * sampleRate)
	osc.releaseSamples = int(release.Seconds() * sampleRate)

	return osc
}

func (osc *Effect) Sine() *Effect {
	osc.wave = sinWave
	return osc
}

func (osc *Effect) Triangle() *Effect {
	osc.wave = triangleWave
	return osc
}

func (osc *Effect) Square() *Effect {
	osc.wave = squareWave
	return osc
}

func (osc *Effect) Sawtooth() *Effect {
	osc.wave = sawtoothWave
	return osc
}

func (osc *Effect) Noise() *Effect {
	osc.wave = noiseWave
	return osc
}

func (osc *Effect) Clone() *Effect {
	clone := new(Effect)
	*clone = *osc

	return clone
}

func (osc *Effect) Play() {
	clone := osc.Clone()

	clone.time = 0
	clone.isPlaying = true

	clone.player = context.NewPlayer(clone)
	clone.player.SetBufferSize(bufferSamples * 4)
	clone.player.Play()

	incrementSoundsPlaying()
}

func (osc *Effect) Read(buf []byte) (n int, err error) {
	for i := 0; i < len(buf); i += 4 {
		if !osc.isPlaying {
			decrementPlayingSounds(osc.player)
			return i, io.EOF
		}

		amplitude := osc.currentAmplitude()
		phase := float64(osc.time) * float64(osc.frequency) / sampleRate

		var sample float32
		switch osc.wave {
		case sinWave:
			sample = float32(math.Sin(phase * math.Pi * 2))
		case triangleWave:
			_, fract := math.Modf(phase)
			sample = -float32(math.Round(fract))*2 + 1
		case squareWave:
			_, fract := math.Modf(phase - 0.25)
			sample = -float32(math.Abs(fract)-0.5)*4 - 1
		case sawtoothWave:
			_, fract := math.Modf(phase - 0.5)
			sample = float32(fract)*2 - 1
		case noiseWave:
			sample = rand.Float32()*2 - 1
		}

		binary.LittleEndian.PutUint32(buf[i:], math.Float32bits(sample*amplitude))

		osc.advanceTime()
	}

	return len(buf), nil
}

func (osc *Effect) advanceTime() {
	osc.time++

	if osc.time >= osc.attackSamples+osc.sustainSamples+osc.releaseSamples {
		osc.isPlaying = false
	}
}

func (osc *Effect) currentAmplitude() float32 {
	if osc.time < osc.attackSamples {
		progress := float32(osc.time) / float32(osc.attackSamples)
		return progress * osc.amplitude
	} else if osc.time < osc.attackSamples+osc.sustainSamples {
		return osc.amplitude
	} else if osc.time < osc.attackSamples+osc.sustainSamples+osc.releaseSamples {
		progress := float32(osc.time-osc.attackSamples-osc.sustainSamples) / float32(osc.releaseSamples)
		return osc.amplitude - osc.amplitude*progress
	} else {
		return 0
	}
}
