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
	fileID := ""

	{
		client := otiai10.NewClient(apiKey)
		fileID = firstFileID(ctx, client)
	}

	{ // otiai10
		client := otiai10.NewClient(apiKey)
		fineTune, err := client.CreateFineTune(ctx, otiai10.FineTuneCreateRequestBody{
			TrainingFile: fileID,
		})
		fmt.Printf("otiai10 FineTune: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(fineTune), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		fineTune, err := client.CreateFineTune(ctx, sashabaranov.FineTuneRequest{
			TrainingFile: fileID,
		})
		fmt.Printf("sashabaranov FineTune: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(fineTune), err)
	}
}

func firstFileID(ctx context.Context, client *otiai10.Client) string {
	listFiles, err := client.ListFiles(ctx)
	if err != nil || len(listFiles.Data) == 0 {
		return ""
	}
	return listFiles.Data[0].ID
}
