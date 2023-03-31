package main

import (
	"context"
	"fmt"

	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		// not supported: deprecated
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		engines, err := client.ListEngines(ctx)
		fmt.Printf("sashabaranov ListEngines: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(engines), err)
	}
}
