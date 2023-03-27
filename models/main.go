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
		models, err := client.ListModels(ctx)
		fmt.Printf("otiai10 ListModels: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(models), err)

		model, err := client.RetrieveModel(ctx, models.Data[0].ID)
		fmt.Printf("otiai10 RetrieveModel: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(model), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		models, err := client.ListModels(ctx)
		fmt.Printf("sashabaranov ListModels: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(models), err)
	}
}
