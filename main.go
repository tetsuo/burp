package main

import (
	"flag"
	"log"
	"net/http"
)

var addrFlag = flag.String("addr", "localhost:9042", "host and port to bind the server to")

func main() {
	mux := http.NewServeMux()

	serverAddr := *addrFlag

	log.Printf("listening on %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, corsHandler(mux)))
}
