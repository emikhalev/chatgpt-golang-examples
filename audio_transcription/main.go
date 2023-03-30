package main

import (
	"context"
	_ "embed"
	"fmt"

	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

const filePath = "assets/this-doesnt-smell-87209.mp3"

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		// not supported
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		imageEdit, err := client.CreateTranscription(ctx, sashabaranov.AudioRequest{
			Model:    "whisper-1",
			FilePath: filePath,
		})
		fmt.Printf("sashabaranov AudioTranscription: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(imageEdit), err)
	}
}
