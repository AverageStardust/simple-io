package sound

import "github.com/ebitengine/oto/v3"

var context *oto.Context

func init() {
	op := &oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	var readyChan chan struct{}
	var err error
	context, readyChan, err = oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}

	<-readyChan
}
