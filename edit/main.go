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
		completion, err := client.CreateEdit(ctx, otiai10.EditCreateRequestBody{
			Model:       "text-davinci-edit-001",
			Input:       "What day of the wek is it?",
			Instruction: "Fix the spelling mistakes",
		})
		fmt.Printf("otiai10 Edits: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		model := "text-davinci-edit-001"
		completion, err := client.Edits(ctx, sashabaranov.EditsRequest{
			Model:       &model,
			Input:       "What day of the wek is it?",
			Instruction: "Fix the spelling mistakes",
		})
		fmt.Printf("sashabaranov Edits: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}
}
