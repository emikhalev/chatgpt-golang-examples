package main

import (
	"context"
	"fmt"

	otiai10 "github.com/otiai10/openaigo"
	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		client := otiai10.NewClient(apiKey)
		completion, err := client.CreateImage(ctx, otiai10.ImageGenerationRequestBody{
			Prompt: "A cute baby sea otter",
			N:      2,
			Size:   "1024x1024",
		})
		fmt.Printf("otiai10 CreateImage: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completion, err := client.CreateImage(ctx, sashabaranov.ImageRequest{
			Prompt: "A cute baby sea otter",
			N:      2,
			Size:   "1024x1024",
		})
		fmt.Printf("sashabaranov CreateImage: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}
}
