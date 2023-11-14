package main

import (
	"context"
	"fmt"

	goopenai "github.com/sashabaranov/go-openai"

	"github.com/eqtlab/lib/core"
	"github.com/eqtlab/lib/openai"
)

func main() {
	openaiClient := goopenai.NewClient("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")

	systemMsg := core.SystemMessage("You are a helpful assistant that translates %s to %s").Bind("English", "French")
	pipe := openai.TextPipe(openaiClient, systemMsg)

	boundUserMsg := core.UserMessage("%s").Bind("I love programming.")

	resultMsg, err := pipe(context.Background(), boundUserMsg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resultMsg.Bytes()))
}
