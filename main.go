package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	anthropicoption "github.com/anthropics/anthropic-sdk-go/option"
	"github.com/openai/openai-go/v2"
	openaioption "github.com/openai/openai-go/v2/option"
)

var addrFlag = flag.String("addr", "localhost:9042", "host and port to bind the server to")

func main() {
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	anthApiKey := os.Getenv("ANTHROPIC_API_KEY")

	if anthApiKey == "" && openaiApiKey == "" {
		log.Fatal("you must set either the OPENAI_API_KEY or the ANTHROPIC_API_KEY environment variable")
	}

	var ac *anthropic.Client
	if anthApiKey != "" {
		c := anthropic.NewClient(anthropicoption.WithAPIKey(anthApiKey))
		ac = &c
		log.Print("enabling Anthropic models")
	}

	var oc *openai.Client
	if openaiApiKey != "" {
		c := openai.NewClient(openaioption.WithAPIKey(openaiApiKey))
		oc = &c
		log.Print("enabling OpenAI models")
	}

	server := &Server{
		wkr: NewWorker(oc, ac),
	}

	mux := http.NewServeMux()

	server.Install(mux)

	serverAddr := *addrFlag

	log.Printf("listening on %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, corsHandler(mux)))
}
