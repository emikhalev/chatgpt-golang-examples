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
		engineID := firstEngineID(ctx, client)

		engineInfo, err := client.GetEngine(ctx, engineID)
		fmt.Printf("sashabaranov GetEngine: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(engineInfo), err)
	}
}

func firstEngineID(ctx context.Context, client *sashabaranov.Client) string {
	list, err := client.ListEngines(ctx)
	if err != nil || len(list.Engines) == 0 {
		return ""
	}
	return list.Engines[0].ID
}
