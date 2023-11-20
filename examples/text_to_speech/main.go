package main

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/eqtlab/lib/core"
	"github.com/eqtlab/lib/openai"
)

func main() {
	input := core.NewUserMessage(`
		One does not simply walk into Mordor.
		Its black gates are guarded by more than just Orcs.
		There is evil there that does not sleep, and the Great Eye is ever watchful.
	`)

	msg, err := openai.New(os.Getenv("OPENAI_API_KEY")).
		TextToSpeech(openai.TextToSpeechParams{
			Model:          "tts-1",
			ResponseFormat: "mp3",
			Speed:          1,
			Voice:          "alloy",
		}).
		Execute(context.Background(), input)

	if err != nil {
		panic(err)
	}

	if err := saveToDisk(msg); err != nil {
		panic(err)
	}
}

func saveToDisk(msg core.Message) error {
	file, err := os.Create("example.mp3")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(msg.Bytes())
	if err != nil {
		return err
	}

	return nil
}