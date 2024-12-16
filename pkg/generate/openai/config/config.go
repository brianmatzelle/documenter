package config

import (
	"os"
)

const OPENAI_MODEL = "gpt-4o"
const OPENAI_SYSTEM_PROMPT = "You are an expert at writing markdown documentation for merge requests."
const OPENAI_PRE_PROMPT_SINGLE = "Generate markdown documentation for this merge request, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it:\n\n"

// const OPENAI_PRE_PROMPT_MULTI = "Generate markdown documentation for the following merge requests, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it. Your are documenting an 'Initiative', so think about the big picture, but still remain concise. This is for company visibility more than technical documentation, so make it look good. Include an 'estimated time to complete the merge request' for each one:\n\n"
const OPENAI_PRE_PROMPT_MULTI = "Generate markdown documentation for the following merge requests, put a TLDR section at the top, explaining why this code exists, then a Key Features section after it. Your are documenting one or more 'Initiative(s)', so think about the big picture, but still remain concise. If there are multiple initiatives, make sure to include the project names when mentioning changes. This is for company visibility more than technical documentation, so make it look good:\n\n"

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
