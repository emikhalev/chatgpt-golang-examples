package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"

	otiai10 "github.com/otiai10/openaigo"
	sashabaranov "github.com/sashabaranov/go-openai"

	"github.com/emikhalev/chatgpt-golang-examples/utils"
)

//go:embed assets/generated-image-1-2.png
var imageFileV1 []byte

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		imageFile := bytes.NewBuffer(imageFileV1)
		client := otiai10.NewClient(apiKey)
		imageVariations, err := client.CreateImageVariation(ctx, otiai10.ImageVariationRequestBody{
			Image: imageFile,
			N:     2,
			Size:  "1024x1024",
		})

		fmt.Printf("otiai10 CreateImageVariation: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(imageVariations), err)
	}

	{ // sashabaranov
		imageFile, err := os.Open("assets/generated-image-2-1.png")
		if err != nil {
			fmt.Printf("sashabaranov EditImage Load file and/or mask error(s): %s, %s\n", err)
		} else {
			client := sashabaranov.NewClient(apiKey)
			imageVariations, err := client.CreateVariImage(ctx, sashabaranov.ImageVariRequest{
				Image: imageFile,
				N:     2,
				Size:  "1024x1024",
			})
			fmt.Printf("sashabaranov EditImage: %v \n (err: %+v)\n\n",
				utils.SPrintStruct(imageVariations), err)
		}
	}
}
