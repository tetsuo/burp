package main

import (
	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/openai/openai-go/v2"
)

type ChatModel string

const (
	// Claude chat models

	ChatModelClaude3_7SonnetLatest      ChatModel = ChatModel(anthropic.ModelClaude3_7SonnetLatest)
	ChatModelClaude3_7Sonnet20250219    ChatModel = ChatModel(anthropic.ModelClaude3_7Sonnet20250219)
	ChatModelClaude3_5HaikuLatest       ChatModel = ChatModel(anthropic.ModelClaude3_5HaikuLatest)
	ChatModelClaude3_5Haiku20241022     ChatModel = ChatModel(anthropic.ModelClaude3_5Haiku20241022)
	ChatModelClaudeSonnet4_20250514     ChatModel = ChatModel(anthropic.ModelClaudeSonnet4_20250514)
	ChatModelClaudeSonnet4_0            ChatModel = ChatModel(anthropic.ModelClaudeSonnet4_0)
	ChatModelClaude4Sonnet20250514      ChatModel = ChatModel(anthropic.ModelClaude4Sonnet20250514)
	ChatModelClaude3_5SonnetLatest      ChatModel = ChatModel(anthropic.ModelClaude3_5SonnetLatest)
	ChatModelClaude3_5Sonnet20241022    ChatModel = ChatModel(anthropic.ModelClaude3_5Sonnet20241022)
	ChatModelClaude_3_5_Sonnet_20240620 ChatModel = ChatModel(anthropic.ModelClaude_3_5_Sonnet_20240620)
	ChatModelClaudeOpus4_0              ChatModel = ChatModel(anthropic.ModelClaudeOpus4_0)
	ChatModelClaudeOpus4_20250514       ChatModel = ChatModel(anthropic.ModelClaudeOpus4_20250514)
	ChatModelClaude4Opus20250514        ChatModel = ChatModel(anthropic.ModelClaude4Opus20250514)
	ChatModelClaudeOpus4_1_20250805     ChatModel = ChatModel(anthropic.ModelClaudeOpus4_1_20250805)
	// Deprecated: Will reach end-of-life on January 5th, 2026. Please migrate to a
	// newer model. Visit
	// https://docs.anthropic.com/en/docs/resources/model-deprecations for more
	// information.
	ChatModelClaude3OpusLatest       ChatModel = ChatModel(anthropic.ModelClaude3OpusLatest)
	ChatModelClaude_3_Opus_20240229  ChatModel = ChatModel(anthropic.ModelClaude_3_Opus_20240229)
	ChatModelClaude_3_Haiku_20240307 ChatModel = ChatModel(anthropic.ModelClaude_3_Haiku_20240307)

	// OpenAI chat models

	ChatModelOpenAIGPT5                             ChatModel = ChatModel(openai.ChatModelGPT5)
	ChatModelOpenAIGPT5Mini                         ChatModel = ChatModel(openai.ChatModelGPT5Mini)
	ChatModelOpenAIGPT5Nano                         ChatModel = ChatModel(openai.ChatModelGPT5Nano)
	ChatModelOpenAIGPT5_2025_08_07                  ChatModel = ChatModel(openai.ChatModelGPT5_2025_08_07)
	ChatModelOpenAIGPT5Mini2025_08_07               ChatModel = ChatModel(openai.ChatModelGPT5Mini2025_08_07)
	ChatModelOpenAIGPT5Nano2025_08_07               ChatModel = ChatModel(openai.ChatModelGPT5Nano2025_08_07)
	ChatModelOpenAIGPT5ChatLatest                   ChatModel = ChatModel(openai.ChatModelGPT5ChatLatest)
	ChatModelOpenAIGPT4_1                           ChatModel = ChatModel(openai.ChatModelGPT4_1)
	ChatModelOpenAIGPT4_1Mini                       ChatModel = ChatModel(openai.ChatModelGPT4_1Mini)
	ChatModelOpenAIGPT4_1Nano                       ChatModel = ChatModel(openai.ChatModelGPT4_1Nano)
	ChatModelOpenAIGPT4_1_2025_04_14                ChatModel = ChatModel(openai.ChatModelGPT4_1_2025_04_14)
	ChatModelOpenAIGPT4_1Mini2025_04_14             ChatModel = ChatModel(openai.ChatModelGPT4_1Mini2025_04_14)
	ChatModelOpenAIGPT4_1Nano2025_04_14             ChatModel = ChatModel(openai.ChatModelGPT4_1Nano2025_04_14)
	ChatModelOpenAIO4Mini                           ChatModel = ChatModel(openai.ChatModelO4Mini)
	ChatModelOpenAIO4Mini2025_04_16                 ChatModel = ChatModel(openai.ChatModelO4Mini2025_04_16)
	ChatModelOpenAIO3                               ChatModel = ChatModel(openai.ChatModelO3)
	ChatModelOpenAIO3_2025_04_16                    ChatModel = ChatModel(openai.ChatModelO3_2025_04_16)
	ChatModelOpenAIO3Mini                           ChatModel = ChatModel(openai.ChatModelO3Mini)
	ChatModelOpenAIO3Mini2025_01_31                 ChatModel = ChatModel(openai.ChatModelO3Mini2025_01_31)
	ChatModelOpenAIO1                               ChatModel = ChatModel(openai.ChatModelO1)
	ChatModelOpenAIO1_2024_12_17                    ChatModel = ChatModel(openai.ChatModelO1_2024_12_17)
	ChatModelOpenAIO1Preview                        ChatModel = ChatModel(openai.ChatModelO1Preview)
	ChatModelOpenAIO1Preview2024_09_12              ChatModel = ChatModel(openai.ChatModelO1Preview2024_09_12)
	ChatModelOpenAIO1Mini                           ChatModel = ChatModel(openai.ChatModelO1Mini)
	ChatModelOpenAIO1Mini2024_09_12                 ChatModel = ChatModel(openai.ChatModelO1Mini2024_09_12)
	ChatModelOpenAIGPT4o                            ChatModel = ChatModel(openai.ChatModelGPT4o)
	ChatModelOpenAIGPT4o2024_11_20                  ChatModel = ChatModel(openai.ChatModelGPT4o2024_11_20)
	ChatModelOpenAIGPT4o2024_08_06                  ChatModel = ChatModel(openai.ChatModelGPT4o2024_08_06)
	ChatModelOpenAIGPT4o2024_05_13                  ChatModel = ChatModel(openai.ChatModelGPT4o2024_05_13)
	ChatModelOpenAIGPT4oAudioPreview                ChatModel = ChatModel(openai.ChatModelGPT4oAudioPreview)
	ChatModelOpenAIGPT4oAudioPreview2024_10_01      ChatModel = ChatModel(openai.ChatModelGPT4oAudioPreview2024_10_01)
	ChatModelOpenAIGPT4oAudioPreview2024_12_17      ChatModel = ChatModel(openai.ChatModelGPT4oAudioPreview2024_12_17)
	ChatModelOpenAIGPT4oAudioPreview2025_06_03      ChatModel = ChatModel(openai.ChatModelGPT4oAudioPreview2025_06_03)
	ChatModelOpenAIGPT4oMiniAudioPreview            ChatModel = ChatModel(openai.ChatModelGPT4oMiniAudioPreview)
	ChatModelOpenAIGPT4oMiniAudioPreview2024_12_17  ChatModel = ChatModel(openai.ChatModelGPT4oMiniAudioPreview2024_12_17)
	ChatModelOpenAIGPT4oSearchPreview               ChatModel = ChatModel(openai.ChatModelGPT4oSearchPreview)
	ChatModelOpenAIGPT4oMiniSearchPreview           ChatModel = ChatModel(openai.ChatModelGPT4oMiniSearchPreview)
	ChatModelOpenAIGPT4oSearchPreview2025_03_11     ChatModel = ChatModel(openai.ChatModelGPT4oSearchPreview2025_03_11)
	ChatModelOpenAIGPT4oMiniSearchPreview2025_03_11 ChatModel = ChatModel(openai.ChatModelGPT4oMiniSearchPreview2025_03_11)
	ChatModelOpenAIChatgpt4oLatest                  ChatModel = ChatModel(openai.ChatModelChatgpt4oLatest)
	ChatModelOpenAICodexMiniLatest                  ChatModel = ChatModel(openai.ChatModelCodexMiniLatest)
	ChatModelOpenAIGPT4oMini                        ChatModel = ChatModel(openai.ChatModelGPT4oMini)
	ChatModelOpenAIGPT4oMini2024_07_18              ChatModel = ChatModel(openai.ChatModelGPT4oMini2024_07_18)
	ChatModelOpenAIGPT4Turbo                        ChatModel = ChatModel(openai.ChatModelGPT4Turbo)
	ChatModelOpenAIGPT4Turbo2024_04_09              ChatModel = ChatModel(openai.ChatModelGPT4Turbo2024_04_09)
	ChatModelOpenAIGPT4_0125Preview                 ChatModel = ChatModel(openai.ChatModelGPT4_0125Preview)
	ChatModelOpenAIGPT4TurboPreview                 ChatModel = ChatModel(openai.ChatModelGPT4TurboPreview)
	ChatModelOpenAIGPT4_1106Preview                 ChatModel = ChatModel(openai.ChatModelGPT4_1106Preview)
	ChatModelOpenAIGPT4VisionPreview                ChatModel = ChatModel(openai.ChatModelGPT4VisionPreview)
	ChatModelOpenAIGPT4                             ChatModel = ChatModel(openai.ChatModelGPT4)
	ChatModelOpenAIGPT4_0314                        ChatModel = ChatModel(openai.ChatModelGPT4_0314)
	ChatModelOpenAIGPT4_0613                        ChatModel = ChatModel(openai.ChatModelGPT4_0613)
	ChatModelOpenAIGPT4_32k                         ChatModel = ChatModel(openai.ChatModelGPT4_32k)
	ChatModelOpenAIGPT4_32k0314                     ChatModel = ChatModel(openai.ChatModelGPT4_32k0314)
	ChatModelOpenAIGPT4_32k0613                     ChatModel = ChatModel(openai.ChatModelGPT4_32k0613)
	ChatModelOpenAIGPT3_5Turbo                      ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo)
	ChatModelOpenAIGPT3_5Turbo16k                   ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo16k)
	ChatModelOpenAIGPT3_5Turbo0301                  ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo0301)
	ChatModelOpenAIGPT3_5Turbo0613                  ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo0613)
	ChatModelOpenAIGPT3_5Turbo1106                  ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo1106)
	ChatModelOpenAIGPT3_5Turbo0125                  ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo0125)
	ChatModelOpenAIGPT3_5Turbo16k0613               ChatModel = ChatModel(openai.ChatModelGPT3_5Turbo16k0613)
)

// ChatProvider identifies which API to use for a given model.
type ChatProvider uint8

const (
	ChatProviderOpenAI ChatProvider = iota + 1
	ChatProviderAnthropic
)

// providerFor is the registry of supported chat models -> provider.
var providerFor = map[ChatModel]ChatProvider{
	ChatModelClaude3_7SonnetLatest:      ChatProviderAnthropic,
	ChatModelClaude3_7Sonnet20250219:    ChatProviderAnthropic,
	ChatModelClaude3_5HaikuLatest:       ChatProviderAnthropic,
	ChatModelClaude3_5Haiku20241022:     ChatProviderAnthropic,
	ChatModelClaudeSonnet4_20250514:     ChatProviderAnthropic,
	ChatModelClaudeSonnet4_0:            ChatProviderAnthropic,
	ChatModelClaude4Sonnet20250514:      ChatProviderAnthropic,
	ChatModelClaude3_5SonnetLatest:      ChatProviderAnthropic,
	ChatModelClaude3_5Sonnet20241022:    ChatProviderAnthropic,
	ChatModelClaude_3_5_Sonnet_20240620: ChatProviderAnthropic,
	ChatModelClaudeOpus4_0:              ChatProviderAnthropic,
	ChatModelClaudeOpus4_20250514:       ChatProviderAnthropic,
	ChatModelClaude4Opus20250514:        ChatProviderAnthropic,
	ChatModelClaudeOpus4_1_20250805:     ChatProviderAnthropic,
	ChatModelClaude3OpusLatest:          ChatProviderAnthropic,
	ChatModelClaude_3_Opus_20240229:     ChatProviderAnthropic,
	ChatModelClaude_3_Haiku_20240307:    ChatProviderAnthropic,

	ChatModelOpenAIGPT5:                             ChatProviderOpenAI,
	ChatModelOpenAIGPT5Mini:                         ChatProviderOpenAI,
	ChatModelOpenAIGPT5Nano:                         ChatProviderOpenAI,
	ChatModelOpenAIGPT5_2025_08_07:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT5Mini2025_08_07:               ChatProviderOpenAI,
	ChatModelOpenAIGPT5Nano2025_08_07:               ChatProviderOpenAI,
	ChatModelOpenAIGPT5ChatLatest:                   ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1:                           ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1Mini:                       ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1Nano:                       ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1_2025_04_14:                ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1Mini2025_04_14:             ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1Nano2025_04_14:             ChatProviderOpenAI,
	ChatModelOpenAIO4Mini:                           ChatProviderOpenAI,
	ChatModelOpenAIO4Mini2025_04_16:                 ChatProviderOpenAI,
	ChatModelOpenAIO3:                               ChatProviderOpenAI,
	ChatModelOpenAIO3_2025_04_16:                    ChatProviderOpenAI,
	ChatModelOpenAIO3Mini:                           ChatProviderOpenAI,
	ChatModelOpenAIO3Mini2025_01_31:                 ChatProviderOpenAI,
	ChatModelOpenAIO1:                               ChatProviderOpenAI,
	ChatModelOpenAIO1_2024_12_17:                    ChatProviderOpenAI,
	ChatModelOpenAIO1Preview:                        ChatProviderOpenAI,
	ChatModelOpenAIO1Preview2024_09_12:              ChatProviderOpenAI,
	ChatModelOpenAIO1Mini:                           ChatProviderOpenAI,
	ChatModelOpenAIO1Mini2024_09_12:                 ChatProviderOpenAI,
	ChatModelOpenAIGPT4o:                            ChatProviderOpenAI,
	ChatModelOpenAIGPT4o2024_11_20:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT4o2024_08_06:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT4o2024_05_13:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT4oAudioPreview:                ChatProviderOpenAI,
	ChatModelOpenAIGPT4oAudioPreview2024_10_01:      ChatProviderOpenAI,
	ChatModelOpenAIGPT4oAudioPreview2024_12_17:      ChatProviderOpenAI,
	ChatModelOpenAIGPT4oAudioPreview2025_06_03:      ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMiniAudioPreview:            ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMiniAudioPreview2024_12_17:  ChatProviderOpenAI,
	ChatModelOpenAIGPT4oSearchPreview:               ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMiniSearchPreview:           ChatProviderOpenAI,
	ChatModelOpenAIGPT4oSearchPreview2025_03_11:     ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMiniSearchPreview2025_03_11: ChatProviderOpenAI,
	ChatModelOpenAIChatgpt4oLatest:                  ChatProviderOpenAI,
	ChatModelOpenAICodexMiniLatest:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMini:                        ChatProviderOpenAI,
	ChatModelOpenAIGPT4oMini2024_07_18:              ChatProviderOpenAI,
	ChatModelOpenAIGPT4Turbo:                        ChatProviderOpenAI,
	ChatModelOpenAIGPT4Turbo2024_04_09:              ChatProviderOpenAI,
	ChatModelOpenAIGPT4_0125Preview:                 ChatProviderOpenAI,
	ChatModelOpenAIGPT4TurboPreview:                 ChatProviderOpenAI,
	ChatModelOpenAIGPT4_1106Preview:                 ChatProviderOpenAI,
	ChatModelOpenAIGPT4VisionPreview:                ChatProviderOpenAI,
	ChatModelOpenAIGPT4:                             ChatProviderOpenAI,
	ChatModelOpenAIGPT4_0314:                        ChatProviderOpenAI,
	ChatModelOpenAIGPT4_0613:                        ChatProviderOpenAI,
	ChatModelOpenAIGPT4_32k:                         ChatProviderOpenAI,
	ChatModelOpenAIGPT4_32k0314:                     ChatProviderOpenAI,
	ChatModelOpenAIGPT4_32k0613:                     ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo:                      ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo16k:                   ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo0301:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo0613:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo1106:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo0125:                  ChatProviderOpenAI,
	ChatModelOpenAIGPT3_5Turbo16k0613:               ChatProviderOpenAI,
}

var (
	// FallbackOpenAIChatModel is used when model not specified.
	FallbackOpenAIChatModel ChatModel = ChatModelOpenAIGPT5Nano
	// FallbackAnthropicChatModel is used when model not specified.
	FallbackAnthropicChatModel ChatModel = ChatModelClaude3_5HaikuLatest
)
