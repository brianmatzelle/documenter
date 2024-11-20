package services

import (
	"documenter/models/requests"
	"fmt"
	"log"

	"documenter/pkg/generate"
	"documenter/pkg/gitlab"
)

func GenerateDocService(request requests.GenDocRequest) (string, error) {
	log.Println("Starting document generation process...")

	mrInfo, err := gitlab.GetMrInfo(request.MrLink, request.GitlabToken)
	if err != nil {
		log.Printf("Failed to fetch MR info: %v", err)
		return "", fmt.Errorf("failed to fetch merge request info: %w", err)
	}
	log.Println("Successfully fetched MR info")

	switch request.Model {
	case "ollama":
		log.Println("Starting document generation with Ollama...")
		docStr, err := generate.GenerateDocOllama(mrInfo)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document")
		return docStr, nil
	case "openai":
		log.Println("Starting document generation with OpenAI...")
		docStr, err := generate.GenerateDocOpenAI(mrInfo)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document")
		return docStr, nil
	default:
		return "", fmt.Errorf("invalid model: %s", request.Model)
	}
}
