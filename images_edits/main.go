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

//go:embed assets/generated-image-1-1.png
var imageFileV1 []byte

//go:embed assets/circle-mask.png
var imageCircleMask []byte

func main() {
	ctx := context.Background()
	apiKey := utils.APIKey()

	{ // otiai10
		imageFile := bytes.NewBuffer(imageFileV1)
		imageMask := bytes.NewBuffer(imageCircleMask)
		client := otiai10.NewClient(apiKey)
		imageEdit, err := client.EditImage(ctx, otiai10.ImageEditRequestBody{
			Image:  imageFile,
			Mask:   imageMask,
			Prompt: "A cute baby sea otter",
			N:      2,
			Size:   "1024x1024",
		})

		fmt.Printf("otiai10 EditImage: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(imageEdit), err)
	}

	{ // sashabaranov
		imageFile, err1 := os.Open("assets/generated-image-2-1.png")
		imageMask, err2 := os.Open("assets/circle-mask.png")
		if err1 != nil || err2 != nil {
			fmt.Printf("sashabaranov EditImage Load file and/or mask error(s): %s, %s\n", err1, err2)
		} else {
			client := sashabaranov.NewClient(apiKey)
			imageEdit, err := client.CreateEditImage(ctx, sashabaranov.ImageEditRequest{
				Image:  imageFile,
				Mask:   imageMask,
				Prompt: "A cute baby sea otter",
				N:      2,
				Size:   "1024x1024",
			})
			fmt.Printf("sashabaranov EditImage: %v \n (err: %+v)\n\n",
				utils.SPrintStruct(imageEdit), err)
		}
	}
}
