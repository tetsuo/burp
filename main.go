package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

var addrFlag = flag.String("addr", "localhost:9042", "host and port to bind the server to")

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENAI_API_KEY is not set")
		os.Exit(1)
	}

	server := &Server{
		wkr: NewWorker(openai.NewClient(option.WithAPIKey(apiKey))),
	}

	mux := http.NewServeMux()

	server.Install(mux)

	serverAddr := *addrFlag

	log.Printf("listening on %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, corsHandler(mux)))
}
