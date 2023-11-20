package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	goopenai "github.com/sashabaranov/go-openai"

	"github.com/eqtlab/lib/core"
	"github.com/eqtlab/lib/openai"
)

func main() {
	factory := openai.New(os.Getenv("OPENAI_API_KEY"))

	resultMsg, err := factory.
		TextToText(openai.TextToTextParams{Model: goopenai.GPT3Dot5Turbo}).
		WithOptions(
			core.WithPrompt(
				"You are a helpful assistant that translates %s to %s",
				"English", "French",
			),
		).
		Execute(
			context.Background(),
			core.NewUserMessage("%s").Bind("I love programming."),
		)

	if err != nil {
		panic(err)
	}

	fmt.Println(resultMsg)
}
