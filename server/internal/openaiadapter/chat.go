package openaiadapter

import (
	"context"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/openai/openai-go/v3"
)

// Takes user prompt and return assistant response along with tokens used
func Chat(ctx context.Context, systemMessage string, history []models.ChatMessage, userMessage string) (responseMessage string, tokensUsed int64, error error) {
	msgs := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemMessage),
	}

	for _, msg := range history {
		switch msg.Role {
		case "user":
			msgs = append(msgs, openai.UserMessage(msg.Message))
		case "assistant":
			msgs = append(msgs, openai.AssistantMessage(msg.Message))
		}
	}

	msgs = append(msgs, openai.UserMessage(userMessage))

	chatParams := openai.ChatCompletionNewParams{
		Messages: msgs,
		Model:    openai.ChatModelGPT4_1Mini,
	}

	chatCompletion, chatErr := openaiClient.Chat.Completions.New(ctx, chatParams)
	if chatErr != nil {
		return "", 0, chatErr
	}

	tokens := chatCompletion.Usage.TotalTokens

	return chatCompletion.Choices[0].Message.Content, tokens, nil
}

// Generates title and summary for a chat session
func GenerateTitleAndSummary(ctx context.Context, messages []models.ChatMessage) (chatTitle string, chatSummary string, error error) {

	titleSystemPrompt := `You are a title generator. Analyze the conversation and create a concise, descriptive title that captures the main topic or purpose of the discussion. The title should be 3-6 words long. Output ONLY the title text itself - no quotation marks, no preamble like "Title:", no explanations, and no additional text.`
	msgs := make([]openai.ChatCompletionMessageParamUnion, len(messages))

	for i, msg := range messages {
		switch msg.Role {
		case "user":
			msgs[i] = openai.UserMessage(msg.Message)
		case "assistant":
			msgs[i] = openai.AssistantMessage(msg.Message)
		}
	}

	titleChatParams := openai.ChatCompletionNewParams{
		Messages: append(
			[]openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(titleSystemPrompt),
			}, msgs...,
		),
		Model: openai.ChatModelGPT4_1Mini,
	}

	summarySystemPrompt := `You are a conversation summarizer. Analyze the conversation and provide a brief, informative summary in 4-6 sentences that captures the key topics discussed, main questions asked, and primary outcomes or conclusions. Output ONLY the summary text itself - no preamble like "Summary:" or "This conversation was about", no labels, and no additional commentary.`

	summaryChatParams := openai.ChatCompletionNewParams{
		Messages: append(
			[]openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(summarySystemPrompt),
			}, msgs...,
		),
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
	summ := summaryChatCompletion.Choices[0].Message.Content

	return title, summ, nil
}
