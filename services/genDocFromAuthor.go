package services

import (
	"documenter/config"
	"documenter/models/requests"
	"documenter/pkg/gitlab"
	"fmt"
)

func GenDocFromAuthor(author string, model string, statusChan chan string) (string, error) {
	// Get GitLab token from environment
	gitlabToken := config.GetEnv("GITLAB_TOKEN")
	if gitlabToken == "" {
		return "", fmt.Errorf("GITLAB_TOKEN environment variable is not set")
	}

	mrLinks, err := gitlab.GetMrLinksFromAuthor(author, gitlabToken)
	if err != nil {
		return "", fmt.Errorf("failed to get merge request links: %w", err)
	}

	request := requests.GenDocRequest{
		MrLinks: mrLinks,
		Model:   model,
	}
	docStr, err := GenerateDocService(request, statusChan)
	if err != nil {
		return "", fmt.Errorf("failed to generate document: %w", err)
	}

	return docStr, nil
}
