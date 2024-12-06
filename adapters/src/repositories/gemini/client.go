package main

import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/google/generative-ai-go/genai"
  "google.golang.org/api/option"
)

// uploadToGemini uploads a file to Gemini and returns the URI of the uploaded file.
func uploadToGemini(ctx context.Context, client *genai.Client, path, mimeType string) string {
  file, err := os.Open(path)
  if err != nil {
    log.Fatalf("Error opening file: %v", err)
  }
  defer file.Close()

  options := genai.UploadFileOptions{
    DisplayName: path,
    MIMEType:    mimeType,
  }
  fileData, err := client.UploadFile(ctx, "", file, &options)
  if err != nil {
    log.Fatalf("Error uploading file: %v", err)
  }

  log.Printf("Uploaded file %s as: %s", fileData.DisplayName, fileData.URI)
  return fileData.URI
}

func main() {
  ctx := context.Background()

  apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
  if !ok {
    log.Fatalln("Environment variable GEMINI_API_KEY not set")
  }
// Create a new client with the API key
  client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
  if err != nil {
    log.Fatalf("Error creating client: %v", err)
  }
  defer client.Close()
// Create a new generative model
// the model is Gemini 1.5 Flash
//  this model is used to generate questions for a quiz based on a PDF file
// there are a limit number of tokens  free, please be aware of that. 
// maybe is a good idea use a different api key for tests

  model := client.GenerativeModel("gemini-1.5-flash")

  model.SetTemperature(1)
  model.SetTopK(40)
  model.SetTopP(0.95)
  model.SetMaxOutputTokens(8192)
  model.ResponseMIMEType = "application/json"
  model.SystemInstructions = &genai.Content{
    // The system instructions are the same as the prompt
    Parts: []genai.Part{genai.Text("You are an expert creating quizzes to evaluate knowledge. Your task is to generate 5 questions following this JSON format:\n\n[\n  {\n    \"question\": \"\",\n    \"option1\": \"\",\n    \"option2\": \"\",\n    \"option3\": \"\",\n    \"option4\": \"\",\n    \"correctOption\": \"\"\n  }\n]\n\n# Rules\n\n- The correct option must always be in a random position.\n- No sentence should exceed 50 characters.\n\n# Output Format\n\n- Produce the questions in the specified JSON format without any additional text.\n\n# Notes\n\n- Ensure each question and answer fits within the character limit.\n- Use diverse questions to test broad knowledge on a subject.")}
  }

  // TODO Make these files available on the local file system
  // You may need to update the file paths
  fileURIs := []string{
    uploadToGemini(ctx, client, "FilePath Here.pdf", "application/pdf"),
  }

  session := model.StartChat()
  session.History = []*genai.Content{
    {
      Role: "user",
      Parts: []genai.Part{
        genai.FileData{URI: fileURIs[0]},
       
      },
    }
  }

  mainTopic := "b2b sales, title example "
  message := fmt.Sprintf("Based on this file, generate the questions for the quiz. The main topic is: %s", mainTopic)
  resp, err := session.SendMessage(ctx, genai.Text(
    message,
  ))
  if err != nil {
    log.Fatalf("Error sending message: %v", err)
  }

  for _, part := range resp.Candidates[0].Content.Parts {
    fmt.Printf("%v\n", part)
  }
}