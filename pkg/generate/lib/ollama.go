package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"documenter/pkg/generate/config"
	"documenter/pkg/generate/models/requests"
	"documenter/pkg/generate/models/responses"
)

func TalkToOllama(ollamaReq requests.OllamaRequest) (*responses.OllamaResponse, error) {
	// Marshal request
	body, err := json.Marshal(&ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ollama request: %w", err)
	}

	// Create client and request
	client := http.Client{}
	httpReq, err := http.NewRequest(http.MethodPost, config.OLLAMA_URL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	// Send request
	log.Printf("sending request to ollama at %s", config.OLLAMA_URL)
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ollama: %w", err)
	}
	defer httpResp.Body.Close()

	// Decode response
	var ollamaResp responses.OllamaResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode ollama response: %w", err)
	}

	log.Printf("successfully received response from ollama")
	return &ollamaResp, nil
}
