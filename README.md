# burp

burp is a chat server that connects clients to OpenAI and Anthropic APIs. It provides HTTP endpoints and a browser-based chat frontend.

## Install

Install it with:

```bash
go install github.com/tetsuo/burp
```

## Usage

#### Start the burp server

```bash
 $ burp
 2025/08/18 23:16:55 enabling anthropic models
 2025/08/18 23:16:55 enabling openai models
 2025/08/18 23:16:55 listening on localhost:9042
```

- Defaults to `localhost:9042`.
- Requires one of:
  - `OPENAI_API_KEY`
  - `ANTHROPIC_API_KEY`

#### Send messages

POST to `/ask?id=<channel>&model=<model>` with body text

```
curl --header "Content-Type: text/plain" \
  --request POST \
  --data "tell me a joke" \
  "http://localhost:9042/chat?id=emu&model=claude-3-haiku-20240307&temp=0.75"
```

#### Receive messages

- `/wait?id=<channel>` - long-poll up to 30s
- `/recent?id=<channel>` - fetch message history

#### Web client

Open `/chat?id=<channel>&model=<model>` in a browser

For example, visit [http://localhost:9042/chat?id=warez&model=gpt-5-nano](http://localhost:9042/chat?id=warez&model=claude-3-haiku-20240307)

## Parameters

Forwarded to the provider:

- `temp` - temperature: \[0.0–2.0] OpenAI, \[0.0–1.0] Anthropic
- `max_tokens` - per-model capped maximum (see [provider.go](./provider.go))
- `top_p` - optional nucleus sampling
- `top_k` - Anthropic only

