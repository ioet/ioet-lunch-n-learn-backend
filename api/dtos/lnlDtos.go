package dtos

type LnLCreateIn struct {
	Name             string `json:"name"`
	PresentationDate string `json:"presentationDate"`
}

type LnLUpdateIn struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	PresentationDate string `json:"presentationDate"`
}

type LnLAddAssistantIn struct {
	ID          string `json:"id"`
	AssistantID string `json:"assistantID"`
}

type LnLAddPresenterIn struct {
	ID          string `json:"id"`
	PresenterID string `json:"presenterID"`
}

type LnLGenerateQuizIn struct {
	ID string `json:"id"`
	PDF string `json:"application/pdf"`
	LunchNLearnName string `json:"lunchNLearnName"`
}