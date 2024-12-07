package generate

import (
	"encoding/json"
	"fmt"
	"time"

	"documenter/pkg/generate/ollama"
	"documenter/pkg/generate/openai"

	"github.com/rs/zerolog/log"
)

func GenerateDocOllama(mrInfos []json.RawMessage, model string, statusChan chan string) (string, error) {
	start := time.Now()
	logger := log.With().Str("func", "GenerateDocOllama").Logger()

	statusChan <- fmt.Sprintf("Preparing to use model: %s", model)
	ollamaReq := ollama.BuildOllamaRequest(mrInfos, model)

	logger.Info().
		Str("model", ollamaReq.Model).
		Bool("stream", ollamaReq.Stream).
		Int("messages", len(ollamaReq.Messages)).
		Msg("Sending Ollama request")

	statusChan <- "Sending request to Ollama..."
	ollamaResp, err := ollama.TalkToOllama(ollamaReq, statusChan)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get Ollama response")
		return "", fmt.Errorf("failed to get Ollama response: %w", err)
	}

	logger.Info().
		Dur("duration", time.Since(start)).
		Msg("Generated document completed")

	return ollamaResp.Message.Content, nil
}

func GenerateDocOpenAI(mrInfos []json.RawMessage) (string, error) {
	start := time.Now()
	logger := log.With().Str("func", "GenerateDocOpenAI").Logger()

	logger.Info().Msg("Starting document generation with OpenAI...")

	logger.Info().
		Dur("duration", time.Since(start)).
		Msg("Generated document completed")

	openaiReq := openai.BuildOpenAIRequest(mrInfos)
	openaiResp, err := openai.TalkToOpenAI(openaiReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get OpenAI response")
		return "", fmt.Errorf("failed to get OpenAI response: %w", err)
	}

	if len(openaiResp.Choices) == 0 {
		logger.Error().Msg("OpenAI returned empty choices")
		return "", fmt.Errorf("OpenAI returned no response choices, openaiResp: %v", openaiResp)
	}

	return openaiResp.Choices[0].Message.Content, nil
}
