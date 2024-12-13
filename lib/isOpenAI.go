package lib

import "slices"

func IsOpenAIModel(model string) bool {
	return slices.Contains(SUPPORTED_OPENAI_MODELS, model)
}
