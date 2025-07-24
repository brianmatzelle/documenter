# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Backend (Go)
- **Build**: `go build .` (builds Go binary)
- **Run locally**: `go run main.go` (runs on port 8080)
- **Dependencies**: `go mod tidy` (manage Go modules)

### Frontend (Next.js)
- **Development**: `cd frontend && npm run dev` (runs on port 3000 with Turbopack)
- **Build**: `cd frontend && npm run build`
- **Production**: `cd frontend && npm run start`
- **Lint**: `cd frontend && npm run lint`
- **Install deps**: `cd frontend && npm install`

### Docker Development
- **Start all services**: `./dev.sh` (starts API, frontend, and Ollama via docker-compose.dev.yml)
- **Production**: `docker compose up -d` (uses docker-compose.yml)

## Architecture Overview

This is a full-stack application that generates AI-powered documentation from GitLab merge requests.

### Backend Structure (Go + Gin)
- **Entry point**: `main.go` - starts Gin server on port 8080
- **Routing**: `controllers/router.go` - defines API endpoints (/ping, /generate-doc, /gen-from-author)
- **Controllers**: `controllers/` directory handles HTTP requests
- **Services**: `services/` contains business logic (genDocService.go)
- **Models**: Request/response structs in `models/requests/` and `models/responses/`

### AI Integration Architecture
The system supports two AI providers through a unified interface:

- **OpenAI Integration**: `pkg/generate/openai/` - GPT-4o integration
- **Ollama Integration**: `pkg/generate/ollama/` - Local LLM (llama3.2:3b) 
- **Provider Selection**: Determined by `lib/isOpenAI.go` based on model name
- **Generation Handler**: `pkg/generate/handler.go` orchestrates AI calls

### GitLab Integration
- **MR Fetching**: `pkg/gitlab/handler.go` fetches merge request data
- **Data Translation**: `pkg/gitlab/lib/translate.go` processes GitLab API responses

### Frontend (Next.js + TypeScript)
- **Framework**: Next.js 15 with Turbopack for fast development
- **UI**: React components in `frontend/src/components/`
- **Context**: Document generation state managed in `frontend/src/context/doc-context.tsx`
- **Styling**: Tailwind CSS with custom UI components

### Configuration
AI model settings are centralized in config files:
- **OpenAI**: `pkg/generate/openai/config/config.go` - model (gpt-4o), prompts, API key
- **Ollama**: `pkg/generate/ollama/config/config.go` - model (llama3.2:3b), URL, prompts

### Key Data Flow
1. Frontend sends POST to `/generate-doc` with MR links, GitLab token, and model choice
2. Service fetches MR data from GitLab API (`pkg/gitlab/`)
3. System routes to appropriate AI provider (`pkg/generate/`)
4. AI generates markdown documentation from MR context
5. Documentation returned to frontend for display

### Docker Architecture
Multi-container setup with docker-compose:
- **API container**: Go backend (port 8080)
- **Frontend container**: Next.js app (port 3000) 
- **Ollama container**: Local LLM server (port 11434)
- **Network**: All containers communicate via `ollama-docker` network

## Environment Setup
- Required: `.env` file with `OPENAI_API_KEY` for OpenAI integration
- GitLab token required at runtime for MR access
- Ollama automatically downloads llama3.2:3b model on first use