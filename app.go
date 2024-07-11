package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAaON84YLb-CnrkcAMNhLvPi9hJyFX9j5A"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Use Client.UploadFile to Upload a file to the service.
	// Pass it an io.Reader.
	f, err := os.Open("data/script.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// You can choose a name, or pass the empty string to generate a unique one.
	file, err := client.UploadFile(ctx, "", f, nil)
	if err != nil {
		log.Fatal(err)
	}
	// The return value's URI field should be passed to the model in a FileData part.
	model := client.GenerativeModel("gemini-1.5-pro")

	resp, err := model.GenerateContent(ctx, genai.FileData{URI: file.URI})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
