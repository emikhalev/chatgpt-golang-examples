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
		completion, err := client.Chat(ctx, otiai10.ChatCompletionRequestBody{
			Model: "gpt-3.5-turbo",
			Messages: []otiai10.ChatMessage{
				{
					Role:    "user",
					Content: "Hello!",
				},
			},
		})
		fmt.Printf("otiai10 ChatCompletion: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completion, err := client.CreateChatCompletion(ctx, sashabaranov.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []sashabaranov.ChatCompletionMessage{
				{
					Role:    "user",
					Content: "Hello!",
				},
			},
		})
		fmt.Printf("sashabaranov ChatCompletion: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}
}
