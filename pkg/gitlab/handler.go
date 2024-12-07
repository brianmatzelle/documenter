package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"documenter/pkg/gitlab/lib"
)

func GetMrInfo(mrLink string, gitlabToken string, model string) (json.RawMessage, error) {
	apiLink, err := lib.TranslateMrLinkToApiLink(mrLink, gitlabToken)
	if err != nil {
		log.Printf("Failed to translate MR link to API link: %v", err)
		return nil, fmt.Errorf("translating MR link: %w", err)
	}

	log.Printf("Fetching merge request info from: %s", apiLink)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiLink, nil)
	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("PRIVATE-TOKEN", gitlabToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to execute HTTP request: %v", err)
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-200 status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, fmt.Errorf("reading response: %w", err)
	}

	log.Printf("Successfully retrieved MR info")
	cleanedResponse, err := cleanMrResponse(body, gitlabToken, model)
	if err != nil {
		return nil, fmt.Errorf("cleaning MR response: %w", err)
	}
	return cleanedResponse, nil
}

func cleanMrResponse(mrResponse json.RawMessage, gitlabToken string, model string) (json.RawMessage, error) {
	log.Printf("Starting to clean MR response")

	var mrInfo map[string]interface{}
	err := json.Unmarshal(mrResponse, &mrInfo)
	if err != nil {
		log.Printf("Failed to unmarshal MR response: %v", err)
		return nil, fmt.Errorf("unmarshalling MR response: %w", err)
	}

	// Extract only the essential information
	cleanedInfo := map[string]interface{}{
		"title":         mrInfo["title"],
		"description":   mrInfo["description"],
		"state":         mrInfo["state"],
		"merged_at":     mrInfo["merged_at"],
		"target_branch": mrInfo["target_branch"],
		"source_branch": mrInfo["source_branch"],
		"reviewers":     mrInfo["reviewers"],
	}

	// Add author name if available
	if author, ok := mrInfo["author"].(map[string]interface{}); ok {
		cleanedInfo["author"] = author["name"]
		log.Printf("Added author: %v", author["name"])
	}

	// // Convert project_id to string properly
	// projectID := fmt.Sprintf("%d", int(mrInfo["project_id"].(float64)))

	// Add changes information
	if changes, ok := mrInfo["changes"].([]interface{}); ok {
		log.Printf("Processing %d file changes", len(changes))
		var fileChanges []map[string]interface{}
		for i, change := range changes {
			if changeMap, ok := change.(map[string]interface{}); ok {
				log.Printf("Processing change %d/%d: %s", i+1, len(changes), changeMap["new_path"])
				// sourceCode, err := getFileSourceCode(projectID, changeMap["new_path"].(string), gitlabToken)
				// if err != nil {
				// 	log.Printf("Failed to get source code for %s: %v", changeMap["new_path"], err)
				// 	return nil, fmt.Errorf("getting file source code: %w", err)
				// }
				fileChange := map[string]interface{}{
					"new_path":    changeMap["new_path"],
					"diff":        changeMap["diff"],
					"source_code": "...",
				}
				fileChanges = append(fileChanges, fileChange)
			}
		}
		cleanedInfo["changes"] = fileChanges
	}

	cleanedJSON, err := json.MarshalIndent(cleanedInfo, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal cleaned response: %v", err)
		return nil, fmt.Errorf("marshalling cleaned response: %w", err)
	}

	log.Printf("Successfully cleaned MR response")
	return cleanedJSON, nil
}

// func getFileSourceCode(projectId string, newPath string, gitlabToken string) (string, error) {
// 	log.Printf("Fetching source code for project %s, path: %s", projectId, newPath)

// 	apiLink, err := lib.TranslatePathToApiLink(projectId, newPath, gitlabToken)
// 	if err != nil {
// 		log.Printf("Failed to translate path to API link: %v", err)
// 		return "", fmt.Errorf("translating path to API link: %w", err)
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodGet, apiLink, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("creating request: %w", err)
// 	}
// 	req.Header.Set("PRIVATE-TOKEN", gitlabToken)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Printf("Failed to execute request for %s: %v", newPath, err)
// 		return "", fmt.Errorf("executing request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Printf("Received non-200 status code (%d) for %s", resp.StatusCode, newPath)
// 		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Failed to read response body for %s: %v", newPath, err)
// 		return "", fmt.Errorf("reading response: %w", err)
// 	}
// 	sourceCode := string(body)
// 	log.Printf("Successfully retrieved source code for %s, length: %d", newPath, len(sourceCode))
// 	return sourceCode, nil
// }
