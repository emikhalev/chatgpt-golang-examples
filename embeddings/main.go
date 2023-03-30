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
		embedding, err := client.CreateEmbedding(ctx, otiai10.EmbeddingCreateRequestBody{
			Model: "text-embedding-ada-002",
			Input: []string{"The food was delicious and the waiter..."},
		})
		fmt.Printf("otiai10 Embedding: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(embedding), err)
	}

	{ // sashabaranov
		client := sashabaranov.NewClient(apiKey)
		embedding, err := client.CreateEmbeddings(ctx, sashabaranov.EmbeddingRequest{
			Model: sashabaranov.AdaEmbeddingV2,
			Input: []string{"What day of the wek is it?"},
		})
		fmt.Printf("sashabaranov Embedding: %v \n (err: %+v)\n\n",
			utils.SPrintStruct(embedding), err)
	}
}
