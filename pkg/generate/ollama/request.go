package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"documenter/pkg/generate/ollama/config"
	"documenter/pkg/generate/ollama/lib"
	"documenter/pkg/generate/ollama/models"
)

func TalkToOllama(ollamaReq models.OllamaRequest, statusChan chan string) (*models.OllamaResponse, error) {
	statusChan <- fmt.Sprintf("Checking if model %s exists...", ollamaReq.Model)
	if err := lib.LoadModel(ollamaReq.Model, statusChan); err != nil {
		return nil, fmt.Errorf("failed to load model: %w", err)
	}

	// Marshal request
	body, err := json.Marshal(&ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ollama request: %w", err)
	}

	chatUrl := config.OLLAMA_URL + "/chat"

	// Create client and request
	client := http.Client{}
	httpReq, err := http.NewRequest(http.MethodPost, chatUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	// Send request
	log.Printf("sending request to ollama at %s", chatUrl)
	log.Printf("request body: %s", string(body))
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ollama: %w", err)
	}
	defer httpResp.Body.Close()

	// Decode response
	var ollamaResp models.OllamaResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode ollama response: %w", err)
	}

	log.Printf("successfully received response from ollama")
	return &ollamaResp, nil
}

func BuildOllamaRequest(mrInfos []json.RawMessage, model string) models.OllamaRequest {
	var mrInfosString string
	for _, mrInfo := range mrInfos {
		mrInfosString += string(mrInfo) + "\n"
	}

	var prePrompt string
	if len(mrInfos) == 1 {
		prePrompt = config.OLLAMA_PRE_PROMPT_SINGLE
	} else {
		prePrompt = config.OLLAMA_PRE_PROMPT_MULTI
	}

	return models.OllamaRequest{
		Model:  model,
		Stream: false,
		Messages: []models.Message{
			{
				Role:    "system",
				Content: config.OLLAMA_SYSTEM_PROMPT,
			},
			{
				Role:    "user",
				Content: prePrompt + mrInfosString,
			},
		},
	}
}
