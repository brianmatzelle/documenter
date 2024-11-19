package generate

import (
	"encoding/json"
	"fmt"
	"time"

	"documenter/pkg/generate/config"
	"documenter/pkg/generate/lib"
	"documenter/pkg/generate/models"
	"documenter/pkg/generate/models/requests"
)

func GenerateDoc(mrInfo json.RawMessage) (string, error) {
	start := time.Now()
	ollamaReq := requests.OllamaRequest{
		Model: config.OLLAMA_MODEL,
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
		Stream: false,
	}
	fmt.Printf("Sending Ollama request:\n  Model: %s\n  Stream: %v\n  Messages: %d\n",
		ollamaReq.Model, ollamaReq.Stream, len(ollamaReq.Messages))

	ollamaResp, err := lib.TalkToOllama(ollamaReq)

	fmt.Println("Got ollama response")
	if err != nil {
		fmt.Println("Error getting ollama response: ", err)
		return "", err
	}
	fmt.Printf("Generated document completed in %v\n", time.Since(start))
	return ollamaResp.Message.Content, nil
}
