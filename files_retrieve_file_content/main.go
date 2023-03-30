package main

import (
	"context"
	"fmt"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
	otiai10 "github.com/otiai10/openaigo"
)

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		client := otiai10.NewClient(apiKey)
		fileID := firstFileID(ctx, client)
		retrieveFileContent, err := client.RetrieveFileContent(ctx, fileID)
		fmt.Printf("otiai10 RetrieveFileContent: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(retrieveFileContent), err)
	}

	{ // sashabaranov
		// not implemented
	}
}

func firstFileID(ctx context.Context, client *otiai10.Client) string {
	listFiles, err := client.ListFiles(ctx)
	if err != nil || len(listFiles.Data) == 0 {
		return ""
	}
	return listFiles.Data[0].ID
}
