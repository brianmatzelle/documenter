package config

// TODO: Make this configurable
// const OLLAMA_URL = "http://localhost:11434/api" // For local testing
const OLLAMA_URL = "http://ollama:11434/api" // For docker-compose

const OLLAMA_MODEL = "llama3.1:8b" // TODO: Make this configurable
const OLLAMA_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OLLAMA_PRE_PROMPT_SINGLE = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"
const OLLAMA_PRE_PROMPT_MULTI = "Generate markdown documentation for the following merge requests, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it. Your are documenting an 'Initiative', so think about the big picture, but still remain concise. This is for company visibility more than technical documentation, so make it look good. Include an 'estimated time to complete the merge request' for each one:\n\n"
