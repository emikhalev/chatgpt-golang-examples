package main

import (
	"context"
	_ "embed"
	"fmt"

	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

const filePath = "assets/life_is_beautiful_italiano-108264.mp3"

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		// not supported
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		imageEdit, err := client.CreateTranslation(ctx, sashabaranov.AudioRequest{
			Model:    "whisper-1",
			FilePath: filePath,
		})
		fmt.Printf("sashabaranov AudioTranslation: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(imageEdit), err)
	}
}
