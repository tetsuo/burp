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
	ChatModelClaude3_7SonnetLatest:      ChatProviderAnthropic, // active, alias to latest Claude 3.7 Sonnet
	ChatModelClaude3_7Sonnet20250219:    ChatProviderAnthropic, // active, Claude 3.7 Sonnet snapshot (2025-02-19)
	ChatModelClaude3_5HaikuLatest:       ChatProviderAnthropic, // active, alias to latest Claude 3.5 Haiku
	ChatModelClaude3_5Haiku20241022:     ChatProviderAnthropic, // active, Claude 3.5 Haiku snapshot (2024-10-22)
	ChatModelClaudeSonnet4_20250514:     ChatProviderAnthropic, // active, Claude 4.0 Sonnet snapshot (2025-05-14)
	ChatModelClaudeSonnet4_0:            ChatProviderAnthropic, // active, Claude 4.0 Sonnet base
	ChatModelClaude4Sonnet20250514:      ChatProviderAnthropic, // active, alt identifier for Claude 4.0 Sonnet (2025-05-14)
	ChatModelClaude3_5SonnetLatest:      ChatProviderAnthropic, // active, alias to latest Claude 3.5 Sonnet
	ChatModelClaude3_5Sonnet20241022:    ChatProviderAnthropic, // active, Claude 3.5 Sonnet snapshot (2024-10-22)
	ChatModelClaude_3_5_Sonnet_20240620: ChatProviderAnthropic, // active, Claude 3.5 Sonnet snapshot (2024-06-20)
	ChatModelClaudeOpus4_0:              ChatProviderAnthropic, // active, Claude 4.0 Opus base
	ChatModelClaudeOpus4_20250514:       ChatProviderAnthropic, // active, Claude 4.0 Opus snapshot (2025-05-14)
	ChatModelClaude4Opus20250514:        ChatProviderAnthropic, // active, alt identifier for Claude 4.0 Opus (2025-05-14)
	ChatModelClaudeOpus4_1_20250805:     ChatProviderAnthropic, // active, Claude 4.1 Opus snapshot (2025-08-05)
	ChatModelClaude3OpusLatest:          ChatProviderAnthropic, // deprecated, alias for Claude 3 Opus (replaced by Claude 4 Opus)
	ChatModelClaude_3_Opus_20240229:     ChatProviderAnthropic, // deprecated, Claude 3 Opus snapshot (2024-02-29)
	ChatModelClaude_3_Haiku_20240307:    ChatProviderAnthropic, // deprecated, Claude 3 Haiku snapshot (2024-03-07)

	ChatModelOpenAIGPT5:                             ChatProviderOpenAI, // active, current flagship (released Aug 2025)
	ChatModelOpenAIGPT5Mini:                         ChatProviderOpenAI, // active, GPT-5 Mini tier
	ChatModelOpenAIGPT5Nano:                         ChatProviderOpenAI, // active, GPT-5 Nano tier
	ChatModelOpenAIGPT5_2025_08_07:                  ChatProviderOpenAI, // active, GPT-5 snapshot (2025-08-07)
	ChatModelOpenAIGPT5Mini2025_08_07:               ChatProviderOpenAI, // active, GPT-5 Mini snapshot (2025-08-07)
	ChatModelOpenAIGPT5Nano2025_08_07:               ChatProviderOpenAI, // active, GPT-5 Nano snapshot (2025-08-07)
	ChatModelOpenAIGPT5ChatLatest:                   ChatProviderOpenAI, // active, alias to latest GPT-5 chat
	ChatModelOpenAIGPT4_1:                           ChatProviderOpenAI, // active, GPT-4.1 family
	ChatModelOpenAIGPT4_1Mini:                       ChatProviderOpenAI, // active, GPT-4.1 Mini
	ChatModelOpenAIGPT4_1Nano:                       ChatProviderOpenAI, // active, GPT-4.1 Nano
	ChatModelOpenAIGPT4_1_2025_04_14:                ChatProviderOpenAI, // active, GPT-4.1 snapshot (2025-04-14)
	ChatModelOpenAIGPT4_1Mini2025_04_14:             ChatProviderOpenAI, // active, GPT-4.1 Mini snapshot
	ChatModelOpenAIGPT4_1Nano2025_04_14:             ChatProviderOpenAI, // active, GPT-4.1 Nano snapshot
	ChatModelOpenAIO4Mini:                           ChatProviderOpenAI, // active in API, pulled from ChatGPT UI after GPT-5 launch
	ChatModelOpenAIO4Mini2025_04_16:                 ChatProviderOpenAI, // active, O4 Mini snapshot (2025-04-16)
	ChatModelOpenAIO3:                               ChatProviderOpenAI, // active, O3
	ChatModelOpenAIO3_2025_04_16:                    ChatProviderOpenAI, // active, O3 snapshot (2025-04-16)
	ChatModelOpenAIO3Mini:                           ChatProviderOpenAI, // active, O3 Mini
	ChatModelOpenAIO3Mini2025_01_31:                 ChatProviderOpenAI, // active, O3 Mini snapshot (2025-01-31)
	ChatModelOpenAIO1:                               ChatProviderOpenAI, // active, O1
	ChatModelOpenAIO1_2024_12_17:                    ChatProviderOpenAI, // active, O1 snapshot (2024-12-17)
	ChatModelOpenAIO1Preview:                        ChatProviderOpenAI, // deprecated, removed Jul 2025
	ChatModelOpenAIO1Preview2024_09_12:              ChatProviderOpenAI, // deprecated, removed Jul 2025
	ChatModelOpenAIO1Mini:                           ChatProviderOpenAI, // deprecated, removal Oct 2025
	ChatModelOpenAIO1Mini2024_09_12:                 ChatProviderOpenAI, // deprecated, removal Oct 2025
	ChatModelOpenAIGPT4o:                            ChatProviderOpenAI, // active in API, pulled from ChatGPT UI after GPT-5 launch
	ChatModelOpenAIGPT4o2024_11_20:                  ChatProviderOpenAI, // active, GPT-4o snapshot (2024-11-20)
	ChatModelOpenAIGPT4o2024_08_06:                  ChatProviderOpenAI, // active, GPT-4o snapshot (2024-08-06)
	ChatModelOpenAIGPT4o2024_05_13:                  ChatProviderOpenAI, // active, GPT-4o snapshot (2024-05-13)
	ChatModelOpenAIGPT4oAudioPreview:                ChatProviderOpenAI, // active alias, but older 2024-10-01 snapshot deprecated
	ChatModelOpenAIGPT4oAudioPreview2024_10_01:      ChatProviderOpenAI, // deprecated, audio-preview snapshot (2024-10-01)
	ChatModelOpenAIGPT4oAudioPreview2024_12_17:      ChatProviderOpenAI, // active, audio-preview snapshot (2024-12-17)
	ChatModelOpenAIGPT4oAudioPreview2025_06_03:      ChatProviderOpenAI, // active, audio-preview snapshot (2025-06-03)
	ChatModelOpenAIGPT4oMiniAudioPreview:            ChatProviderOpenAI, // active, GPT-4o Mini audio-preview alias
	ChatModelOpenAIGPT4oMiniAudioPreview2024_12_17:  ChatProviderOpenAI, // active, GPT-4o Mini audio-preview snapshot
	ChatModelOpenAIGPT4oSearchPreview:               ChatProviderOpenAI, // active (preview model, subject to change)
	ChatModelOpenAIGPT4oMiniSearchPreview:           ChatProviderOpenAI, // active (preview model, subject to change)
	ChatModelOpenAIGPT4oSearchPreview2025_03_11:     ChatProviderOpenAI, // active, search-preview snapshot (2025-03-11)
	ChatModelOpenAIGPT4oMiniSearchPreview2025_03_11: ChatProviderOpenAI, // active, mini search-preview snapshot (2025-03-11)
	ChatModelOpenAIChatgpt4oLatest:                  ChatProviderOpenAI, // alias, not an API model (maps to latest GPT-4o)
	ChatModelOpenAICodexMiniLatest:                  ChatProviderOpenAI, // deprecated, Codex family retired Mar 2023
	ChatModelOpenAIGPT4oMini:                        ChatProviderOpenAI, // active, GPT-4o Mini
	ChatModelOpenAIGPT4oMini2024_07_18:              ChatProviderOpenAI, // active, GPT-4o Mini snapshot (2024-07-18)
	ChatModelOpenAIGPT4Turbo:                        ChatProviderOpenAI, // active, GPT-4 Turbo
	ChatModelOpenAIGPT4Turbo2024_04_09:              ChatProviderOpenAI, // active, GA GPT-4 Turbo snapshot (2024-04-09)
	ChatModelOpenAIGPT4_0125Preview:                 ChatProviderOpenAI, // deprecated, replaced by GPT-4 Turbo GA
	ChatModelOpenAIGPT4TurboPreview:                 ChatProviderOpenAI, // deprecated, replaced by GPT-4 Turbo GA
	ChatModelOpenAIGPT4_1106Preview:                 ChatProviderOpenAI, // deprecated, replaced by GPT-4 Turbo GA
	ChatModelOpenAIGPT4VisionPreview:                ChatProviderOpenAI, // deprecated, superseded by GPT-4o multimodal
	ChatModelOpenAIGPT4:                             ChatProviderOpenAI, // active, GPT-4 base family
	ChatModelOpenAIGPT4_0314:                        ChatProviderOpenAI, // deprecated, retired mid-2024
	ChatModelOpenAIGPT4_0613:                        ChatProviderOpenAI, // deprecated, retired mid-2024
	ChatModelOpenAIGPT4_32k:                         ChatProviderOpenAI, // deprecated, 32k family retired mid-2024
	ChatModelOpenAIGPT4_32k0314:                     ChatProviderOpenAI, // deprecated, retired mid-2024
	ChatModelOpenAIGPT4_32k0613:                     ChatProviderOpenAI, // deprecated, retired mid-2024
	ChatModelOpenAIGPT3_5Turbo:                      ChatProviderOpenAI, // active, GPT-3.5 Turbo family
	ChatModelOpenAIGPT3_5Turbo16k:                   ChatProviderOpenAI, // deprecated, replaced when 16k became default
	ChatModelOpenAIGPT3_5Turbo0301:                  ChatProviderOpenAI, // deprecated, retired 2024
	ChatModelOpenAIGPT3_5Turbo0613:                  ChatProviderOpenAI, // deprecated, retired 2024
	ChatModelOpenAIGPT3_5Turbo1106:                  ChatProviderOpenAI, // active, GPT-3.5 Turbo snapshot (2023-11-06)
	ChatModelOpenAIGPT3_5Turbo0125:                  ChatProviderOpenAI, // active, GPT-3.5 Turbo snapshot (2024-01-25)
	ChatModelOpenAIGPT3_5Turbo16k0613:               ChatProviderOpenAI, // deprecated, retired 2024
}

var modelMaxOutputTokens = map[ChatModel]int64{
	// Claude 3.7 Sonnet
	ChatModelClaude3_7SonnetLatest:   64_000,
	ChatModelClaude3_7Sonnet20250219: 64_000,

	// Claude 3.5
	ChatModelClaude3_5HaikuLatest:       8_192,
	ChatModelClaude3_5Haiku20241022:     8_192,
	ChatModelClaude3_5SonnetLatest:      8_000,
	ChatModelClaude3_5Sonnet20241022:    8_000,
	ChatModelClaude_3_5_Sonnet_20240620: 8_000,

	// Claude 4.0 / 4.1
	ChatModelClaudeSonnet4_20250514: 64_000,
	ChatModelClaudeSonnet4_0:        64_000,
	ChatModelClaude4Sonnet20250514:  64_000,
	ChatModelClaudeOpus4_0:          32_000,
	ChatModelClaudeOpus4_20250514:   32_000,
	ChatModelClaude4Opus20250514:    32_000,
	ChatModelClaudeOpus4_1_20250805: 32_000,

	// Claude 3 (deprecated)
	ChatModelClaude3OpusLatest:       4_000,
	ChatModelClaude_3_Opus_20240229:  4_000,
	ChatModelClaude_3_Haiku_20240307: 4_096,

	// GPT-5 family (assume same as GPT-4.1 until doc updates)
	ChatModelOpenAIGPT5:               32_000,
	ChatModelOpenAIGPT5Mini:           32_000,
	ChatModelOpenAIGPT5Nano:           32_000,
	ChatModelOpenAIGPT5_2025_08_07:    32_000,
	ChatModelOpenAIGPT5Mini2025_08_07: 32_000,
	ChatModelOpenAIGPT5Nano2025_08_07: 32_000,
	ChatModelOpenAIGPT5ChatLatest:     32_000,

	// GPT-4.1 family
	ChatModelOpenAIGPT4_1:               32_000,
	ChatModelOpenAIGPT4_1Mini:           32_000,
	ChatModelOpenAIGPT4_1Nano:           32_000,
	ChatModelOpenAIGPT4_1_2025_04_14:    32_000,
	ChatModelOpenAIGPT4_1Mini2025_04_14: 32_000,
	ChatModelOpenAIGPT4_1Nano2025_04_14: 32_000,

	// O-series
	ChatModelOpenAIO4Mini:              100_000,
	ChatModelOpenAIO4Mini2025_04_16:    100_000,
	ChatModelOpenAIO3:                  100_000,
	ChatModelOpenAIO3_2025_04_16:       100_000,
	ChatModelOpenAIO3Mini:              100_000,
	ChatModelOpenAIO3Mini2025_01_31:    100_000,
	ChatModelOpenAIO1:                  100_000,
	ChatModelOpenAIO1_2024_12_17:       100_000,
	ChatModelOpenAIO1Preview:           32_000,
	ChatModelOpenAIO1Preview2024_09_12: 32_000,
	ChatModelOpenAIO1Mini:              65_536,
	ChatModelOpenAIO1Mini2024_09_12:    65_536,

	// GPT-4o + Mini + Turbo
	ChatModelOpenAIGPT4o:                            16_000,
	ChatModelOpenAIGPT4o2024_11_20:                  16_000,
	ChatModelOpenAIGPT4o2024_08_06:                  16_000,
	ChatModelOpenAIGPT4o2024_05_13:                  16_000,
	ChatModelOpenAIGPT4oAudioPreview:                16_000,
	ChatModelOpenAIGPT4oAudioPreview2024_10_01:      16_000,
	ChatModelOpenAIGPT4oAudioPreview2024_12_17:      16_000,
	ChatModelOpenAIGPT4oAudioPreview2025_06_03:      16_000,
	ChatModelOpenAIGPT4oMiniAudioPreview:            16_000,
	ChatModelOpenAIGPT4oMiniAudioPreview2024_12_17:  16_000,
	ChatModelOpenAIGPT4oSearchPreview:               16_000,
	ChatModelOpenAIGPT4oMiniSearchPreview:           16_000,
	ChatModelOpenAIGPT4oSearchPreview2025_03_11:     16_000,
	ChatModelOpenAIGPT4oMiniSearchPreview2025_03_11: 16_000,
	ChatModelOpenAIChatgpt4oLatest:                  16_000,
	ChatModelOpenAIGPT4oMini:                        16_000,
	ChatModelOpenAIGPT4oMini2024_07_18:              16_000,
	ChatModelOpenAIGPT4Turbo:                        16_000,
	ChatModelOpenAIGPT4Turbo2024_04_09:              16_000,
	ChatModelOpenAIGPT4_0125Preview:                 16_000,
	ChatModelOpenAIGPT4TurboPreview:                 16_000,
	ChatModelOpenAIGPT4_1106Preview:                 16_000,
	ChatModelOpenAIGPT4VisionPreview:                16_000,

	// GPT-4 base
	ChatModelOpenAIGPT4:         16_000,
	ChatModelOpenAIGPT4_0314:    16_000,
	ChatModelOpenAIGPT4_0613:    16_000,
	ChatModelOpenAIGPT4_32k:     16_000, // OpenAI docs: 32k context, but max output capped at 16k
	ChatModelOpenAIGPT4_32k0314: 16_000,
	ChatModelOpenAIGPT4_32k0613: 16_000,

	// GPT-3.5
	ChatModelOpenAIGPT3_5Turbo:        4_000,
	ChatModelOpenAIGPT3_5Turbo16k:     4_000, // 16k ctx, still 4k out
	ChatModelOpenAIGPT3_5Turbo0301:    4_000,
	ChatModelOpenAIGPT3_5Turbo0613:    4_000,
	ChatModelOpenAIGPT3_5Turbo1106:    4_000,
	ChatModelOpenAIGPT3_5Turbo0125:    4_000,
	ChatModelOpenAIGPT3_5Turbo16k0613: 4_000,
}

var modelMaxInputChars = map[ChatModel]int64{
	// Anthropic ~200k ctx
	ChatModelClaude3_7SonnetLatest:   680_000,
	ChatModelClaude3_7Sonnet20250219: 680_000,
	ChatModelClaudeSonnet4_20250514:  680_000,
	ChatModelClaudeSonnet4_0:         680_000,
	ChatModelClaude4Sonnet20250514:   680_000,
	ChatModelClaudeOpus4_0:           680_000,
	ChatModelClaudeOpus4_20250514:    680_000,
	ChatModelClaude4Opus20250514:     680_000,
	ChatModelClaudeOpus4_1_20250805:  680_000,
	ChatModelClaude3OpusLatest:       680_000,
	ChatModelClaude_3_Opus_20240229:  680_000,

	// Anthropic Haiku ~64k ctx
	ChatModelClaude3_5HaikuLatest:    250_000,
	ChatModelClaude3_5Haiku20241022:  250_000,
	ChatModelClaude_3_Haiku_20240307: 250_000,

	// Anthropic 3.5 Sonnet ~200k ctx
	ChatModelClaude3_5SonnetLatest:      680_000,
	ChatModelClaude3_5Sonnet20241022:    680_000,
	ChatModelClaude_3_5_Sonnet_20240620: 680_000,

	// OpenAI 128k ctx families
	ChatModelOpenAIGPT5:                 392_000,
	ChatModelOpenAIGPT5Mini:             392_000,
	ChatModelOpenAIGPT5Nano:             392_000,
	ChatModelOpenAIGPT5_2025_08_07:      392_000,
	ChatModelOpenAIGPT5Mini2025_08_07:   392_000,
	ChatModelOpenAIGPT5Nano2025_08_07:   392_000,
	ChatModelOpenAIGPT5ChatLatest:       392_000,
	ChatModelOpenAIGPT4_1:               392_000,
	ChatModelOpenAIGPT4_1Mini:           392_000,
	ChatModelOpenAIGPT4_1Nano:           392_000,
	ChatModelOpenAIGPT4_1_2025_04_14:    392_000,
	ChatModelOpenAIGPT4_1Mini2025_04_14: 392_000,
	ChatModelOpenAIGPT4_1Nano2025_04_14: 392_000,
	ChatModelOpenAIO4Mini:               392_000,
	ChatModelOpenAIO4Mini2025_04_16:     392_000,
	ChatModelOpenAIO3:                   392_000,
	ChatModelOpenAIO3_2025_04_16:        392_000,
	ChatModelOpenAIO3Mini:               392_000,
	ChatModelOpenAIO3Mini2025_01_31:     392_000,
	ChatModelOpenAIO1:                   392_000,
	ChatModelOpenAIO1_2024_12_17:        392_000,
	ChatModelOpenAIO1Preview:            392_000,
	ChatModelOpenAIO1Preview2024_09_12:  392_000,
	ChatModelOpenAIO1Mini:               392_000,
	ChatModelOpenAIO1Mini2024_09_12:     392_000,

	// GPT-4o family & Turbo
	ChatModelOpenAIGPT4o:                            392_000,
	ChatModelOpenAIGPT4o2024_11_20:                  392_000,
	ChatModelOpenAIGPT4o2024_08_06:                  392_000,
	ChatModelOpenAIGPT4o2024_05_13:                  392_000,
	ChatModelOpenAIGPT4oAudioPreview:                392_000,
	ChatModelOpenAIGPT4oAudioPreview2024_10_01:      392_000,
	ChatModelOpenAIGPT4oAudioPreview2024_12_17:      392_000,
	ChatModelOpenAIGPT4oAudioPreview2025_06_03:      392_000,
	ChatModelOpenAIGPT4oMiniAudioPreview:            392_000,
	ChatModelOpenAIGPT4oMiniAudioPreview2024_12_17:  392_000,
	ChatModelOpenAIGPT4oSearchPreview:               392_000,
	ChatModelOpenAIGPT4oMiniSearchPreview:           392_000,
	ChatModelOpenAIGPT4oSearchPreview2025_03_11:     392_000,
	ChatModelOpenAIGPT4oMiniSearchPreview2025_03_11: 392_000,
	ChatModelOpenAIChatgpt4oLatest:                  392_000,
	ChatModelOpenAIGPT4oMini:                        392_000,
	ChatModelOpenAIGPT4oMini2024_07_18:              392_000,
	ChatModelOpenAIGPT4Turbo:                        392_000,
	ChatModelOpenAIGPT4Turbo2024_04_09:              392_000,
	ChatModelOpenAIGPT4_0125Preview:                 392_000,
	ChatModelOpenAIGPT4TurboPreview:                 392_000,
	ChatModelOpenAIGPT4_1106Preview:                 392_000,
	ChatModelOpenAIGPT4VisionPreview:                392_000,

	// GPT-4 (8k)
	ChatModelOpenAIGPT4:      24_500,
	ChatModelOpenAIGPT4_0314: 24_500,
	ChatModelOpenAIGPT4_0613: 24_500,

	// GPT-4-32k
	ChatModelOpenAIGPT4_32k:     98_000,
	ChatModelOpenAIGPT4_32k0314: 98_000,
	ChatModelOpenAIGPT4_32k0613: 98_000,

	// GPT-3.5 family
	ChatModelOpenAIGPT3_5Turbo:        12_250,
	ChatModelOpenAIGPT3_5Turbo16k:     44_500,
	ChatModelOpenAIGPT3_5Turbo0301:    12_250,
	ChatModelOpenAIGPT3_5Turbo0613:    12_250,
	ChatModelOpenAIGPT3_5Turbo1106:    12_250,
	ChatModelOpenAIGPT3_5Turbo0125:    12_250,
	ChatModelOpenAIGPT3_5Turbo16k0613: 44_500,
}
