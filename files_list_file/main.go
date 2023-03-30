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
		listFiles, err := client.ListFiles(ctx)
		fmt.Printf("otiai10 ListFiles: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(listFiles), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		listFiles, err := client.ListFiles(ctx)
		fmt.Printf("sashabaranov ListFiles: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(listFiles), err)
	}
}
