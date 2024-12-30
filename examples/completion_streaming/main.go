package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	goopenai "github.com/sashabaranov/go-openai"

	"github.com/nbox363/agency"
	"github.com/nbox363/agency/providers/openai"
)

func main() {
	factory := openai.New(openai.Params{Key: os.Getenv("OPENAI_API_KEY")})

	result, err := factory.
		TextToStream(openai.TextToStreamParams{Model: goopenai.GPT3Dot5Turbo}, func(delta string) error {
			fmt.Printf(delta)
			return nil
		}).
		SetPrompt("Write a few sentences about topic").
		Execute(context.Background(), agency.UserMessage("I love programming."))

	if err != nil {
		panic(err)
	}

	fmt.Println("\nFinal result:", string(result.Content))
}
