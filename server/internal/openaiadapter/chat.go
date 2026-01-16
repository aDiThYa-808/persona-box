package openaiadapter

import (
	"context"

	"github.com/openai/openai-go/v3"
)

func Chat(ctx context.Context, systemMessage string, assistantMessage string, userMessage string) (responseMessage string, tokensUsed int64, error error) {
	chatParams := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMessage),
			openai.AssistantMessage(assistantMessage),
			openai.UserMessage(userMessage),
		},
		Model: openai.ChatModelGPT4_1Mini,
	}

	chatCompletion, chatErr := openaiClient.Chat.Completions.New(ctx, chatParams)
	if chatErr != nil {
		return "", 0, chatErr
	}

	tokens := chatCompletion.Usage.TotalTokens

	return chatCompletion.Choices[0].Message.Content, tokens, nil
}

