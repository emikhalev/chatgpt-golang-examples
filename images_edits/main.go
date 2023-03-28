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
		completion, err := client.EditImage(ctx, otiai10.ImageEditRequestBody{
			Image:  "@otter.png",
			Mask:   "@mask.png",
			Prompt: "A cute baby sea otter",
			N:      2,
			Size:   "1024x1024",
		})
		fmt.Printf("otiai10 EditImage: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		completion, err := client.CreateEditImage(ctx, sashabaranov.ImageEditRequest{
			Image:  "",
			Mask:   "",
			Prompt: "A cute baby sea otter",
			N:      2,
			Size:   "1024x1024",
		})
		fmt.Printf("sashabaranov EditImage: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(completion), err)
	}
}
