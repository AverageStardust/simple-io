package sound

import (
	"time"

	"github.com/ebitengine/oto/v3"
)

const sampleRate = 48000
const bufferSamples = 512

var context *oto.Context
var playingSounds = 0
var lastEndedSound *oto.Player

func init() {
	op := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: 1,
		Format:       oto.FormatFloat32LE,
		BufferSize:   time.Second / sampleRate * bufferSamples,
	}

	var ready chan struct{}
	var err error
	context, ready, err = oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext() failed: " + err.Error())
	}

	<-ready
}

func WaitForSoundsToStop() {
	for areSoundsPlaying() {
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Second / sampleRate * bufferSamples)
	time.Sleep(time.Second / sampleRate * bufferSamples)
}

func areSoundsPlaying() bool {
	if playingSounds > 0 {
		return true
	} else if lastEndedSound != nil && lastEndedSound.IsPlaying() {
		// all sound generated, but still playing on hardware
		return true
	} else {
		return false
	}
}

func incrementSoundsPlaying() {
	playingSounds++
}

func decrementPlayingSounds(player *oto.Player) {
	playingSounds--
	lastEndedSound = player
}
