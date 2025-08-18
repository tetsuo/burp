# burp

burp is a chat server that connects clients to OpenAI and Anthropic APIs. It provides HTTP endpoints and a browser-based chat frontend.

Install with:

```bash
go get github.com/tetsuo/burp
```

## Features

- Supports OpenAI and Anthropic models
- Endpoints:

  - `/ask` - send a message
  - `/wait` - long-poll for new messages
  - `/recent` - fetch chat history
  - `/chat` - web client
