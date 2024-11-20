package config

import (
	"os"
)

const OPENAI_MODEL = "gpt-4o"
const OPENAI_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OPENAI_PRE_PROMPT = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
