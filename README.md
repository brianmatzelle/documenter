# Documenter - AI-Powered Documentation Generator for Merge Requests

Documenter is a tool that automatically generates documentation and release notes from merge requests using AI models (OpenAI or Ollama). It simplifies the documentation process by analyzing merge request details and producing structured markdown documentation.

## Features

* Multiple AI Model Support:
* OpenAI (GPT-4) integration
* Ollama (local LLM) support using llama3.2:3b model
* Docker-based Deployment: Easy setup using Docker Compose
* RESTful API: Simple HTTP endpoints for document generation
* Structured Output: Generates markdown documentation with TLDR and Key Features sections

## Demo

![documenter-demo](https://github.com/brianmatzelle/documenter/blob/main/demos/documenter-demo.gif?raw=true)

## Quick Start

1. Create a .env file with your configuration:
   ```
   OPENAI_API_KEY=your_openai_key_here
   GITLAB_TOKEN=your_gitlab_token_here
   ```
2. Start the service using Docker Compose:
   `docker compose up -d`
3. Go to `http://localhost:3000` to use the frontend

### For local development

1. To start all development services, in the root directory run:
   `./dev.sh`
2. Go to `http://localhost:3000` to use the frontend, or `http://localhost:8080` to use the API.

This will start both the API service and Ollama service (for local LLM support).

## API Documentation

### Base URL

http://localhost:8080

### Endpoints

#### Health Check

`GET /ping`

#### Generate Documentation

`GET /generate-doc` (Server-Sent Events)

Query Parameters:
- `mrLinks`: JSON array of merge request URLs (required)
- `model`: AI model to use - "gpt-4o" or "llama3.2" (required)

Example:
```
GET /generate-doc?mrLinks=["https://gitlab.com/your-project/merge_requests/1"]&model=gpt-4o
```

#### Generate Documentation from Author

`GET /gen-from-author` (Server-Sent Events)

Query Parameters:
- `author`: GitLab username (required)
- `model`: AI model to use - "gpt-4o" or "llama3.2" (required)

Example:
```
GET /gen-from-author?author=username&model=gpt-4o
```

**Note**: Both endpoints use Server-Sent Events (SSE) to stream status updates and return the final documentation. The GitLab token is configured as an environment variable on the server side for security.

Response Events:
- `status`: Progress updates during generation
- `complete`: Final generated documentation
- `error`: Error messages if generation fails

Status Codes:

* 200: Success
* 400: Bad Request
* 500: Internal Server Error

## Architecture

### Backend Components

1. API Layer: Built with Go's Gin framework
   - Reference:

```go
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)
	router.GET("/generate-doc", GenerateDocController)
	router.GET("/gen-from-author", GenFromAuthorController)

	return router
}
```

2. Document Generation Service: Supports multiple AI models
   - Reference:

```go
func GenerateDocService(request requests.GenDocRequest, statusChan chan string) (string, error) {
	log.Println("Starting document generation process...")

	// Get GitLab token from environment
	gitlabToken := config.GetEnv("GITLAB_TOKEN")
	if gitlabToken == "" {
		return "", fmt.Errorf("GITLAB_TOKEN environment variable is not set")
	}

	var mrInfos []json.RawMessage
	for _, mrLink := range request.MrLinks {
		mrInfo, err := gitlab.GetMrInfo(mrLink, gitlabToken, request.Model)
		if err != nil {
			log.Printf("Failed to fetch MR info: %v", err)
			return "", fmt.Errorf("failed to fetch merge request info: %w", err)
		}
		log.Println("Successfully fetched MR info for ", mrLink)
		mrInfos = append(mrInfos, mrInfo)
	}

	if lib.IsOpenAIModel(request.Model) {
		statusChan <- fmt.Sprintf("[OpenAI]: %s selected...", request.Model)
		docStr, err := generate.GenerateDocOpenAI(mrInfos)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document")
		return docStr, nil
	} else {
		statusChan <- fmt.Sprintf("[Ollama]: %s selected...", request.Model)
		docStr, err := generate.GenerateDocOllama(mrInfos, request.Model, statusChan)
		if err != nil {
			log.Printf("Failed to generate document: %v", err)
			return "", fmt.Errorf("failed to generate document: %w", err)
		}
		log.Println("Successfully generated document")
		return docStr, nil
	}
}
```

3. AI Integration:
   - OpenAI Integration

```go
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
```

- Ollama Integration

```go
func TalkToOllama(ollamaReq models.OllamaRequest) (*models.OllamaResponse, error) {
	if err := lib.LoadModel(ollamaReq.Model); err != nil {
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
```

### Docker Configuration

The application uses a multi-container setup with Docker Compose:

```yaml
services:
  api:
    image: brianmatzelle/documenter:latest
    ports:
      - 8080:8080
    container_name: api
    depends_on:
      - ollama
    environment:
      - GIN_MODE=release
      - PORT=8080
    env_file:
      - .env
    networks:
      - ollama-docker

  ollama:
    image: ollama/ollama:latest
    ports:
      - 11434:11434 # for testing, remove the :11434 later
    volumes:
      - .:/code
      - ./ollama-volume:/root/.ollama
    container_name: ollama
    # pull_policy: always
    restart: always
    environment:
      - OLLAMA_KEEP_ALIVE=24h
      - OLLAMA_HOST=0.0.0.0
    networks:
      - ollama-docker

networks:
  ollama-docker:
    external: false
```

## Configuration

Adjust these configs to change models, add new ones, etc.

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```
OPENAI_API_KEY=your_openai_api_key_here
GITLAB_TOKEN=your_gitlab_token_here
PORT=8080  # Optional, defaults to 8080
```

### OpenAI Settings

in `pkg/generate/openai/config/config.go`

```go
const OPENAI_MODEL = "gpt-4o"
const OPENAI_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OPENAI_PRE_PROMPT = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
```

### Ollama Settings

in `pkg/generate/ollama/config/config.go`

```go
// const OLLAMA_URL = "http://localhost:11434/api/chat" // For local testing
const OLLAMA_URL = "http://ollama:11434/api" // For docker-compose
const OLLAMA_MODEL = "llama3.2:3b"           // TODO: Make this configurable
const OLLAMA_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OLLAMA_PRE_PROMPT = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"
```

## Example Usage

Using curl with Server-Sent Events:

```bash
# Generate documentation from merge request links
curl -N "http://localhost:8080/generate-doc?mrLinks=[\"https://gitlab.com/your-project/merge_requests/1\"]&model=gpt-4o"

# Generate documentation from author's recent merge requests
curl -N "http://localhost:8080/gen-from-author?author=your-username&model=llama3.2"
```

Using JavaScript (EventSource):

```javascript
const eventSource = new EventSource('/generate-doc?' + new URLSearchParams({
  mrLinks: JSON.stringify(['https://gitlab.com/your-project/merge_requests/1']),
  model: 'gpt-4o'
}));

eventSource.addEventListener('status', (event) => {
  const data = JSON.parse(event.data);
  console.log('Status:', data.message);
});

eventSource.addEventListener('complete', (event) => {
  const data = JSON.parse(event.data);
  console.log('Generated documentation:', data.doc);
  eventSource.close();
});

eventSource.addEventListener('error', (event) => {
  const data = JSON.parse(event.data);
  console.error('Error:', data.error);
  eventSource.close();
});
```

## Notes

- Ensure you have Docker and Docker Compose installed
- The Ollama service may take some time to initially download and load the LLM model
- GitLab token is configured as an environment variable on the server side for security
- GitLab token requires appropriate permissions to access merge request information
- Choose between "gpt-4o" or "llama3.2" as the model parameter based on your needs
- The API uses Server-Sent Events (SSE) for real-time status updates during document generation

## Todo

- [X] Add AI support for
  - [X] OpenAI
  - [X] Ollama (local)
  - [ ] Anthropic
- [X] Parse and clean the merge request data to improve prompt quality
- [X] Move GitLab token to server-side environment variables for security
- [ ] Refine the pre-prompt to produce more desirable documentation results
- [ ] Implement chat functionality so the user can talk with the merge request.
  - [ ] "Adjust X and include Y..."
  - [ ] "This merge request took 2 weeks to complete. Is that reasonable?"
- [ ] Create a second generation option aimed to help Product Managers, giving a less technical overview of the code
- [ ] Support GitHub merge requests
