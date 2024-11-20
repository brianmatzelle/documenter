package lib

import (
	"bytes"
	"documenter/pkg/generate/ollama/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
)

type ModelResponse struct {
	Models []Model `json:"models"`
}

type Model struct {
	Name string `json:"name"`
	// other fields omitted as they're not needed for our use case
}

func LoadModel(model string) error {
	log.Printf("Checking if model %s exists...", model)
	if modelExists(model) {
		log.Printf("Model %s already exists", model)
		return nil
	}

	err := pullModel(model)
	if err != nil {
		return fmt.Errorf("failed to pull model %s: %w", model, err)
	}

	return nil
}

func pullModel(model string) error {
	log.Printf("Pulling model %s...", model)

	url := fmt.Sprintf("%s/pull", config.OLLAMA_URL)
	jsonStr := []byte(fmt.Sprintf(`{"model": "%s"}`, model))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("failed to pull model: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("Response: %s", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to pull model (status %d): %s", resp.StatusCode, string(body))
	}

	log.Printf("Successfully pulled model %s", model)
	return nil
}

func modelExists(model string) bool {
	models, err := getModelList()
	if err != nil {
		log.Printf("Error checking model existence: %v", err)
		return false
	}

	exists := slices.Contains(models, model)
	log.Printf("Model %s exists: %v", model, exists)
	return exists
}

func getModelList() ([]string, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/tags", config.OLLAMA_URL)

	log.Printf("Fetching model list from %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for model list: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch model list: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read model list response: %w", err)
	}

	var modelResp ModelResponse
	if err := json.Unmarshal(body, &modelResp); err != nil {
		return nil, fmt.Errorf("failed to parse model list response: %w", err)
	}

	models := make([]string, len(modelResp.Models))
	for i, model := range modelResp.Models {
		models[i] = model.Name
	}

	log.Printf("Found %d models", len(models))
	return models, nil
}
