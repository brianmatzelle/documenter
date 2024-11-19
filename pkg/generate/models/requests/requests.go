package requests

import ollama "documenter/pkg/generate/models"

type OllamaRequest struct {
	Model    string           `json:"model"`
	Messages []ollama.Message `json:"messages"`
	Stream   bool             `json:"stream"`
}
