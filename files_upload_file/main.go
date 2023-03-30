package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"

	otiai10 "github.com/otiai10/openaigo"
	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

//go:embed assets/completion.jsonl
var completionFileV1 []byte

const filePath = "assets/completion.jsonl"

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		client := otiai10.NewClient(apiKey)
		completionFile := bytes.NewBuffer(completionFileV1)
		uploadFile, err := client.UploadFile(ctx, otiai10.FileUploadRequestBody{
			Purpose: "fine-tune",
			File:    completionFile,
		})
		fmt.Printf("otiai10 UploadFile: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(uploadFile), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		uploadFile, err := client.CreateFile(ctx, sashabaranov.FileRequest{
			Purpose:  "fine-tune",
			FileName: "completion2.jsonl",
			FilePath: filePath,
		})
		fmt.Printf("sashabaranov UploadFile: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(uploadFile), err)
	}
}
