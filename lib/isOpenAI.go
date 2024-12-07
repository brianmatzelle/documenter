package lib

import "slices"

func IsOpenAI(model string) bool {
	return slices.Contains(SUPPORTED_OPENAI_MODELS, model)
}
