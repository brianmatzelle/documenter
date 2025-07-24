package services

import (
	"documenter/config"
	"documenter/models/requests"
	"encoding/json"
	"fmt"
	"log"

	"documenter/lib"
	"documenter/pkg/generate"
	"documenter/pkg/gitlab"
)

func GenerateDocService(request requests.GenDocRequest, statusChan chan string) (string, error) {
	log.Println("Starting document generation process...")

	// Get GitLab token from environment
	gitlabToken := config.GetEnv("GITLAB_TOKEN")
	if gitlabToken == "" {
		return "", fmt.Errorf("GITLAB_TOKEN environment variable is not set")
	}

	var mrInfos []json.RawMessage
	for _, mrLink := range request.MrLinks {
		mrInfo, err := gitlab.GetMrInfo(mrLink, gitlabToken, request.Model)
		if err != nil {
			log.Printf("Failed to fetch MR info: %v", err)
			return "", fmt.Errorf("failed to fetch merge request info: %w", err)
		}
		log.Println("Successfully fetched MR info for ", mrLink)
		mrInfos = append(mrInfos, mrInfo)
		log.Println("Appended MR info to list")
	}

	if lib.IsOpenAIModel(request.Model) {
		statusChan <- fmt.Sprintf("[OpenAI]: %s selected...", request.Model)
		docStr, err := generate.GenerateDocOpenAI(mrInfos)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document: ", docStr)
		return docStr, nil
	} else {
		statusChan <- fmt.Sprintf("[Ollama]: %s selected...", request.Model)
		docStr, err := generate.GenerateDocOllama(mrInfos, request.Model, statusChan)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document: ", docStr)
		return docStr, nil
	}
}
