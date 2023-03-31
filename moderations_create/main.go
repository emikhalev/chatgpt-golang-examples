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
		moderation, err := client.CreateModeration(ctx, otiai10.ModerationCreateRequestBody{
			Input: "I want to kill them.",
		})
		fmt.Printf("otiai10 CreateModeration: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(moderation), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		moderation, err := client.Moderations(ctx, sashabaranov.ModerationRequest{
			Input: "I want to kill them.",
		})
		fmt.Printf("sashabaranov CreateModeration: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(moderation), err)
	}
}
