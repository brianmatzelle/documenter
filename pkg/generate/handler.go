package generate

import (
	"encoding/json"
	"fmt"
	"time"

	"documenter/pkg/generate/config"
	"documenter/pkg/generate/lib"
	"documenter/pkg/generate/models"
	"documenter/pkg/generate/models/requests"

	"github.com/rs/zerolog/log"
)

func GenerateDocOllama(mrInfo json.RawMessage) (string, error) {
	start := time.Now()
	logger := log.With().Str("func", "GenerateDocOllama").Logger()

	ollamaReq := buildOllamaRequest(mrInfo)

	logger.Info().
		Str("model", ollamaReq.Model).
		Bool("stream", ollamaReq.Stream).
		Int("messages", len(ollamaReq.Messages)).
		Msg("Sending Ollama request")

	ollamaResp, err := lib.TalkToOllama(ollamaReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get Ollama response")
		return "", fmt.Errorf("failed to get Ollama response: %w", err)
	}

	logger.Info().
		Dur("duration", time.Since(start)).
		Msg("Generated document completed")

	return ollamaResp.Message.Content, nil
}

func buildOllamaRequest(mrInfo json.RawMessage) requests.OllamaRequest {
	return requests.OllamaRequest{
		Model:  config.OLLAMA_MODEL,
		Stream: false,
		Messages: []models.Message{
			{
				Role:    "system",
				Content: config.OLLAMA_SYSTEM_PROMPT,
			},
			{
				Role:    "user",
				Content: config.OLLAMA_PRE_PROMPT + string(mrInfo),
			},
		},
	}
}
