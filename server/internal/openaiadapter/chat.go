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

func GenerateTitleAndSummary(ctx context.Context, messages string) (chatTitle string, chatSummary string, error error) {

	titleSystemPrompt := `You are a title generator. Analyze the conversation and create a concise, descriptive title that captures the main topic or purpose of the discussion. The title should be 3-6 words long. Output ONLY the title text itself - no quotation marks, no preamble like "Title:", no explanations, and no additional text.`

	titleChatParams := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(titleSystemPrompt),
			openai.UserMessage(messages),
		},
		Model: openai.ChatModelGPT4_1Mini,
	}

	summarySystemPrompt := `You are a conversation summarizer. Analyze the conversation and provide a brief, informative summary in 2-4 sentences that captures the key topics discussed, main questions asked, and primary outcomes or conclusions. Output ONLY the summary text itself - no preamble like "Summary:" or "This conversation was about", no labels, and no additional commentary.`

	summaryChatParams := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(summarySystemPrompt),
			openai.UserMessage(messages),
		},
		Model: openai.ChatModelGPT4_1Mini,
	}

	titleChatCompletion, titleChatErr := openaiClient.Chat.Completions.New(ctx, titleChatParams)
	if titleChatErr != nil {
		return "", "", titleChatErr
	}

	summaryChatCompletion, summaryChatErr := openaiClient.Chat.Completions.New(ctx, summaryChatParams)
	if summaryChatErr != nil {
		return "", "", summaryChatErr
	}

	title := titleChatCompletion.Choices[0].Message.Content
	summary := summaryChatCompletion.Choices[0].Message.Content

	return title, summary, nil
}
