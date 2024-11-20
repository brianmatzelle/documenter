package generate

import (
	"encoding/json"
	"fmt"
	"time"

	"documenter/pkg/generate/ollama"

	"github.com/rs/zerolog/log"
)

func GenerateDocOllama(mrInfo json.RawMessage) (string, error) {
	start := time.Now()
	logger := log.With().Str("func", "GenerateDocOllama").Logger()

	ollamaReq := ollama.BuildOllamaRequest(mrInfo)

	logger.Info().
		Str("model", ollamaReq.Model).
		Bool("stream", ollamaReq.Stream).
		Int("messages", len(ollamaReq.Messages)).
		Msg("Sending Ollama request")

	ollamaResp, err := ollama.TalkToOllama(ollamaReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get Ollama response")
		return "", fmt.Errorf("failed to get Ollama response: %w", err)
	}

	logger.Info().
		Dur("duration", time.Since(start)).
		Msg("Generated document completed")

	return ollamaResp.Message.Content, nil
}
