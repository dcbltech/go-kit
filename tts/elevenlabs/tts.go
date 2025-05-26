package elevenlabs

import (
	"context"
	"time"

	"github.com/haguro/elevenlabs-go"
)

const (
	modelID = "eleven_multilingual_v2"

	VoiceSally   = "JZz8TL6s5j2bYyqcdMGd"
	VoiceCassidy = "56AoDkrOh6qfVPDXZ7Pt"
	VoiceMark    = "UgBBYS2sOqTuMpoF3BR0"
)

type TextToSpeech struct {
	client *elevenlabs.Client
}

func Must(apiKey string) *TextToSpeech {
	return &TextToSpeech{
		client: elevenlabs.NewClient(context.Background(), apiKey, 10*time.Minute),
	}
}

func (t *TextToSpeech) Synthesize(text, voice string) (data []byte, err error) {
	return t.client.TextToSpeech(
		voice,
		elevenlabs.TextToSpeechRequest{
			Text:    text,
			ModelID: modelID,
		},
	)
}
