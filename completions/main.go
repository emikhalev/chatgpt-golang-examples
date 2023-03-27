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
		completion, err := client.Completion(ctx, otiai10.CompletionRequestBody{
			Model:       "text-davinci-003",
			Prompt:      []string{"Say this is a test"},
			MaxTokens:   7,
			Temperature: 0,
			TopP:        1,
			N:           1,
			Stream:      false,
			Stop:        []string{"\n"},
		})
		fmt.Printf("otiai10 Completion: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completion, err := client.CreateCompletion(ctx, sashabaranov.CompletionRequest{
			Model:       "text-davinci-003",
			Prompt:      "Say this is a test",
			MaxTokens:   7,
			Temperature: 0,
			TopP:        1,
			N:           1,
			Stream:      false,
			Stop:        []string{"\n"},
		})
		fmt.Printf("sashabaranov Completion: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}
}
