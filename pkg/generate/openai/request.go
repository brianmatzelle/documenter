package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"documenter/pkg/generate/openai/config"
	"documenter/pkg/generate/openai/models"
)

func TalkToOpenAI(openaiReq models.OpenAIRequest) (*models.OpenAIResponse, error) {
	// Marshal request
	body, err := json.Marshal(&openaiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal openai request: %w", err)
	}

	chatUrl := "https://api.openai.com/v1/chat/completions"

	// Create client and request
	client := http.Client{}
	httpReq, err := http.NewRequest(http.MethodPost, chatUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	// Add authorization header
	httpReq.Header.Add("Authorization", "Bearer "+config.OPENAI_API_KEY)
	httpReq.Header.Add("Content-Type", "application/json")

	// Send request
	log.Printf("sending request to OpenAI at %s", chatUrl)
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to openai: %w", err)
	}
	defer httpResp.Body.Close()

	// Decode response
	var openaiResp models.OpenAIResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&openaiResp); err != nil {
		return nil, fmt.Errorf("failed to decode openai response: %w", err)
	}

	log.Printf("successfully received response from OpenAI")
	return &openaiResp, nil
}

func BuildOpenAIRequest(mrInfo json.RawMessage) models.OpenAIRequest {
	return models.OpenAIRequest{
		Model: config.OPENAI_MODEL,
		Messages: []models.Message{
			{
				Role:    "system",
				Content: config.OPENAI_SYSTEM_PROMPT,
			},
			{
				Role:    "user",
				Content: config.OPENAI_PRE_PROMPT + string(mrInfo),
			},
		},
	}
}