package lnlgeneratequiznusecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"github.com/google/generative-ai-go/genai"
)

type GenerateQuizUseCase struct {
	client *genai.Client
}

func NewGenerateQuizUseCase(client *genai.Client) *GenerateQuizUseCase {
	return &GenerateQuizUseCase{
		client: client,
	}
}

func (uc *GenerateQuizUseCase) Execute(ctx context.Context, file multipart.File, fileName, mimeType string) ([]string, error) {
	options := genai.UploadFileOptions{
		DisplayName: fileName,
		MIMEType:    mimeType,
	}

	fileData, err := uc.client.UploadFile(ctx, "", file, &options)
	if err != nil {
		return nil, fmt.Errorf("error uploading file: %w", err)
	}

	// Setup and call model to generate questions
	model := uc.client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.SystemInstructions = &genai.Content{
		Parts: []genai.Part{
			genai.Text(`generate a JSON with 5 questions on a specific topic, using the attached file provided. Considerer the outpur must be like this in JSON FORMAT: [{question,option_1,option_2,option_3,option_4,correct_option}]. Also considerer that the correct option should be in a random position, not always the same.`),
		},
	}

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.FileData{URI: fileData.URI},
				genai.Text("questions about: " + fileName),
			},
		},
	}

	resp, err := session.SendMessage(ctx, genai.Text("Generate questions"))
	if err != nil {
		return nil, fmt.Errorf("error sending message: %w", err)
	}

	questions := []string{}
	for _, part := range resp.Candidates[0].Content.Pparts {
		questions = append(questions, part.String())
	}

	return questions, nil
}
