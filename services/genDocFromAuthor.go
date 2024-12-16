package services

import (
	"documenter/models/requests"
	"documenter/pkg/gitlab"
	"fmt"
)

func GenDocFromAuthor(author string, gitlabToken string, model string, statusChan chan string) (string, error) {
	mrLinks, err := gitlab.GetMrLinksFromAuthor(author, gitlabToken)
	if err != nil {
		return "", fmt.Errorf("failed to get merge request links: %w", err)
	}

	request := requests.GenDocRequest{
		MrLinks:     mrLinks,
		GitlabToken: gitlabToken,
		Model:       model,
	}
	docStr, err := GenerateDocService(request, statusChan)
	if err != nil {
		return "", fmt.Errorf("failed to generate document: %w", err)
	}

	return docStr, nil
}
