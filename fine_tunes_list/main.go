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
		fineTunesList, err := client.ListFineTunes(ctx)
		fmt.Printf("otiai10 ListFineTunes: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(fineTunesList), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		fineTunesList, err := client.ListFineTunes(ctx)
		fmt.Printf("sashabaranov ListFineTunes: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(fineTunesList), err)
	}
}
