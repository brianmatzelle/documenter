package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// https://gitlab.icg360.net/eng/keystone/-/merge_requests/878 -> https://gitlab.icg360.net/api/v4/projects/347/merge_requests/878/changes
func TranslateMrLinkToApiLink(mrLink string, gitlabToken string) (string, error) {
	if mrLink == "" || gitlabToken == "" {
		return "", fmt.Errorf("mrLink and gitlabToken are required")
	}

	projectId, err := getProjectIdFromMrLink(mrLink, gitlabToken)
	if err != nil {
		return "", fmt.Errorf("failed to get project ID: %w", err)
	}
	if projectId == "" {
		return "", fmt.Errorf("project ID not found for link: %s", mrLink)
	}

	mrId := getMrIdFromMrLink(mrLink)
	apiLink := fmt.Sprintf("https://gitlab.icg360.net/api/v4/projects/%s/merge_requests/%s/changes", projectId, mrId)
	fmt.Printf("Generated API link: %s\n", apiLink)
	return apiLink, nil
}

func getProjectIdFromMrLink(mrLink string, gitlabToken string) (string, error) {
	parts := strings.Split(mrLink, "/")
	if len(parts) < 4 {
		return "", fmt.Errorf("invalid merge request link format")
	}

	projectName := parts[3]
	if projectName == "eng" {
		projectName = parts[4]
	}
	fmt.Printf("Searching for project: %s\n", projectName)

	url := fmt.Sprintf("https://gitlab.icg360.net/api/v4/projects?search=%s", projectName)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("PRIVATE-TOKEN", gitlabToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var projects []map[string]interface{}
	if err := json.Unmarshal(body, &projects); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	for _, project := range projects {
		if project["name"] == projectName {
			id := strconv.Itoa(int(project["id"].(float64)))
			fmt.Printf("Found project ID: %s\n", id)
			return id, nil
		}
	}

	return "", fmt.Errorf("project not found: %s", projectName)
}

func getMrIdFromMrLink(mrLink string) string {
	parts := strings.Split(mrLink, "/")
	mrId := parts[len(parts)-1]
	fmt.Printf("Extracted merge request ID: %s\n", mrId)
	return mrId
}