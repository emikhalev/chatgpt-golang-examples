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
		// not supported
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completionStream, err := client.CreateCompletionStream(ctx, sashabaranov.CompletionRequest{
			Model:       "text-davinci-003",
			Prompt:      "Say this is a test",
			MaxTokens:   7,
			Temperature: 0,
			TopP:        1,
			N:           1,
			Stream:      false,
			Stop:        []string{"\n"},
		})

		resp := sashabaranov.CompletionResponse{}
		for {
			resp, err = completionStream.Recv()
			if err == io.EOF {
				fmt.Printf("sashabaranov CreateCompletionStream streamReader.Recv(): EOF\n")
				break
			} else if err != nil {
				fmt.Printf("sashabaranov CreateCompletionStream streamReader.Recv() error: %v\n", err)
				break
			}
			fmt.Printf("sashabaranov CreateCompletionStream streamReader.Recv(): \n%v\n\n", utils.SPrintStruct(resp))
		}
	}
}
