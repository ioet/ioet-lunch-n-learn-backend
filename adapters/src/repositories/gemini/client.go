package GeminiClient

import "net/http"

type GeminiAPIClient struct {
	url string
	apiKey string
}

func NewPeopleAPIClient(httpClient *http.Client) *GeminiAPIClient {
	return &GeminiAPIClient{
		host:       envvar.PeopleAPIUrl(),
		httpClient: httpClient,
		roles:      []string{envvar.PeopleAPIRoleName()},
	}
}



package main

import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/google/generative-ai-go/genai"
  "google.golang.org/api/option"
)

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

  client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
  if err != nil {
    log.Fatalf("Error creating client: %v", err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  model.SetTemperature(1)
  model.SetTopK(64)
  model.SetTopP(0.95)
  model.SetMaxOutputTokens(8192)
  model.ResponseMIMEType = "application/json"
  model.SystemInstructions = &genai.Content{
    Parts: []genai.Part{genai.Text("You are a fun question generator. Your goal is to generate 5 multiple choice questions with the text that the user provides you.\nin JSON format using this example\n\n[\n{question,\noption_1, option_2,option_3,option_4, correct:option1\n}\n]")}
  }

  // model.SafetySettings = Adjust safety settings
  // See https://ai.google.dev/gemini-api/docs/safety-settings

  // TODO Make these files available on the local file system
  // You may need to update the file paths
  fileURIs := []string{
    uploadToGemini(ctx, client, "image_food4.jpeg", "image/jpeg"),
  }

  session := model.StartChat()
  session.History = []*genai.Content{
    {
      Role: "user",
      Parts: []genai.Part{
        genai.FileData{URI: fileURIs[0]},
        genai.Text("Given this image detail the recipe to bake this item in JSON format. Include item names and quantities for the recipe."),
      },
    },
    {
      Role: "model",
      Parts: []genai.Part{
        genai.Text("```json\n{\"recipe\": {\"name\": \"Sweet Potato Fries\", \"ingredients\": [{\"name\": \"sweet potatoes\", \"quantity\": \"1\", \"unit\": \"medium\"}, {\"name\": \"olive oil\", \"quantity\": \"2\", \"unit\": \"tablespoons\"}, {\"name\": \"salt\", \"quantity\": \"1/2\", \"unit\": \"teaspoon\"}, {\"name\": \"black pepper\", \"quantity\": \"1/4\", \"unit\": \"teaspoon\"}, {\"name\": \"garlic powder\", \"quantity\": \"1/4\", \"unit\": \"teaspoon\"}, {\"name\": \"onion powder\", \"quantity\": \"1/4\", \"unit\": \"teaspoon\"}], \"instructions\": [\"Preheat oven to 400 degrees F (200 degrees C).\", \"Cut sweet potatoes into 1/2-inch thick fries.\", \"In a large bowl, toss sweet potato fries with olive oil, salt, pepper, garlic powder, and onion powder.\", \"Spread sweet potato fries in a single layer on a baking sheet.\", \"Bake in preheated oven for 20-25 minutes, or until tender and golden brown.\", \"Serve immediately.\"}}\n\n```"),
      },
    },
  }

  resp, err := session.SendMessage(ctx, genai.Text("INSERT_INPUT_HERE"))
  if err != nil {
    log.Fatalf("Error sending message: %v", err)
  }

  for _, part := range resp.Candidates[0].Content.Parts {
    fmt.Printf("%v\n", part)
  }
}