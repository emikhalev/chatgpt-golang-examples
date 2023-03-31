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
	fineTuneID := ""

	{
		client := otiai10.NewClient(apiKey)
		fineTuneID = firstFineTuneID(ctx, client)
	}

	{ // otiai10
		client := otiai10.NewClient(apiKey)
		deleteFineTune, err := client.DeleteFineTuneModel(ctx, fineTuneID)
		fmt.Printf("otiai10 DeleteFineTuneModel: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(deleteFineTune), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		deleteFineTune, err := client.DeleteFineTune(ctx, fineTuneID)
		fmt.Printf("sashabaranov DeleteFineTuneModel: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(deleteFineTune), err)
	}
}

func firstFineTuneID(ctx context.Context, client *otiai10.Client) string {
	list, err := client.ListFineTunes(ctx)
	if err != nil || len(list.Data) == 0 {
		return ""
	}
	return list.Data[0].ID
}
