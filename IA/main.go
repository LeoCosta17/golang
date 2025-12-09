package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func main() {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "",
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-pro-preview",
		genai.Text("Olá tudo bem? Estou começando a trabalhar em aplicações com IA. Poderia montar um road map para estudo?"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}
