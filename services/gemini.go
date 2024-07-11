package scripts

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
)

func RunDream(text string) genai.Part {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	//apiKey := os.Getenv("")

	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		option.WithAPIKey("AIzaSyAaON84YLb-CnrkcAMNhLvPi9hJyFX9j5A"))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("넌 루나야 루나"),
		},
		Role: "user",
	}

	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("ㅎㅇ 루나 이제부터 넌 날 도와야해"),
			},
			Role: "user",
		},
	}
	resp, err := cs.SendMessage(ctx, genai.Text(text))
	if err != nil {
		log.Fatal(err)
	}
	str := PrintModelResp(resp)
	return str
}

func PrintModelResp(resp *genai.GenerateContentResponse) genai.Part {
	var content genai.Part
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				content = part // genai.Part의 Text를 반환합니다.
			}
		}
	}
	return content
}
