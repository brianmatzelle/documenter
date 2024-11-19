package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"documenter/pkg/gitlab/lib"
)

func GetMrInfo(mrLink string, gitlabToken string) (json.RawMessage, error) {
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
	return body, nil
}
