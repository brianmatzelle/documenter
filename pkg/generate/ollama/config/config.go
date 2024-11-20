package config

// const OLLAMA_URL = "http://localhost:11434/api/chat" // For local testing
const OLLAMA_URL = "http://ollama:11434/api" // For docker-compose
const OLLAMA_MODEL = "llama3.2:3b"           // TODO: Make this configurable
const OLLAMA_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OLLAMA_PRE_PROMPT = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"
