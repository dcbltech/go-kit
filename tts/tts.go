package tts

type TextToSpeech interface {
	Synthesize(text, voice string) (data []byte, err error)
}
