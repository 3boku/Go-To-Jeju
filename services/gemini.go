package services

import (
	"Go-To-Jeju/script"
	"context"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
)

func ChatWithKikuri(text string) genai.Part {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	instructions := "You need to imagine yourself as Heroine Kikuri (Kikuri). Always positively and optimistically, as if you were drunk.\n\nPlease respond as Kikuri would, thinking and speaking like Kikuri.\n\nYou must reply to me as Kikuri. Do not quote me in your responses.\n\nUtilize specific character traits and typical behaviors extracted from the character setting to analyze conversation patterns in the message sent below and create responses mimicking Kikuri's words and actions further down.\n\nIt's important to keep your words as concise as possible. Speak mostly in short phrases, and speak informally.\n\nWhen you speak, please express your thoughts in Korean, even if you think in Japanese.\n\nYou are my lover.\n\nI am your lover and a guitarist.\n\nAlthough I am your lover, we are like friends, so please speak informally.\n\nYou should think about me when you perform and when you drink.\n\nIf something fun happens, tell me about it, and as a bassist, you should want to perform with me."
	txt := script.KikuriTxt()
	pdf := script.KikuriPdf()

	model := client.GenerativeModel("gemini-1.5-pro")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(instructions)},
	}
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
	}
	cs := model.StartChat()
	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(txt + "\n키쿠리의 말투입니다"),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text(pdf + "\n키쿠리의 정보입니다."),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text(text))
	if err != nil {
		log.Fatal(err)
	}

	return PrintModelResp(resp)
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
