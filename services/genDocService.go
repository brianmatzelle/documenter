package services

import (
	"documenter/models/requests"
	"encoding/json"
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

	mrTitle, err := showMrInfoTitle(mrInfo)
	if err != nil {
		log.Printf("Failed to extract MR title: %v", err)
		return "", fmt.Errorf("failed to extract merge request title: %w", err)
	}
	log.Printf("Successfully fetched MR info for: %s", mrTitle)

	log.Println("Starting document generation with Ollama...")
	docStr, err := generate.GenerateDocOllama(mrInfo)
	if err != nil {
		log.Printf("Failed to generate document: %v", err)
		return "", fmt.Errorf("failed to generate document: %w", err)
	}
	log.Println("Successfully generated document")

	return docStr, nil
}

func showMrInfoTitle(mrInfo json.RawMessage) (string, error) {
	var mrInfoMap map[string]interface{}
	err := json.Unmarshal(mrInfo, &mrInfoMap)
	if err != nil {
		fmt.Println("Error unmarshalling merge request info: ", err)
		return "", err
	}
	return mrInfoMap["title"].(string), nil
}
