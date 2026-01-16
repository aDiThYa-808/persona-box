package openaiadapter

import (
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var openaiClient openai.Client

func Init() {
	openaiKey := os.Getenv("OPENAI_SECRET_KEY")
	if openaiKey == "" {
		panic("open ai key not found")
	}

	openaiClient = openai.NewClient(
		option.WithAPIKey(openaiKey),
	)
}
