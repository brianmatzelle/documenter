package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"documenter/pkg/generate/config"
	"documenter/pkg/generate/models/requests"
	"documenter/pkg/generate/models/responses"
)

func TalkToOllama(ollamaReq requests.OllamaRequest) (*responses.OllamaResponse, error) {
	fmt.Println("Marshalling ollama request")
	body, err := json.Marshal(&ollamaReq)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	fmt.Println("Creating http request for ollama")
	httpReq, err := http.NewRequest(http.MethodPost, config.OLLAMA_URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	fmt.Println("Sending http request to ollama, url: ", config.OLLAMA_URL)
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	fmt.Println("Decoding ollama response")
	ollamaResp := responses.OllamaResponse{}
	err = json.NewDecoder(httpResp.Body).Decode(&ollamaResp)
	return &ollamaResp, err
}
