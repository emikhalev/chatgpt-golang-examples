package main

import (
	"context"
	"fmt"
	"io"

	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		// not supported, but here is a pull request to add such support:
		// https://github.com/otiai10/openaigo/pull/14
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completionStream, err := client.CreateChatCompletionStream(ctx, sashabaranov.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []sashabaranov.ChatCompletionMessage{
				{
					Role:    "user",
					Content: "Hello!",
				},
			},
		})

		resp := sashabaranov.ChatCompletionStreamResponse{}
		for {
			resp, err = completionStream.Recv()
			if err == io.EOF {
				fmt.Printf("sashabaranov ChatCompletionResponse streamReader.Recv(): EOF\n")
				break
			} else if err != nil {
				fmt.Printf("sashabaranov ChatCompletionResponse streamReader.Recv() error: %v\n", err)
				break
			}
			fmt.Printf("sashabaranov ChatCompletionResponse streamReader.Recv(): \n%v\n\n", utils.SPrintStruct(resp))
		}
	}
}
