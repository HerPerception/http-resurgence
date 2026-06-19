package main

import (
	"log"
	"net/http"
)

// func main() {
func main() {
	// The main mux — receives all requests
	mainMux := http.NewServeMux()

	// A sub-mux — handles only /api/* routes
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	// Mount apiMux under /api/ — StripPrefix removes "/api"
	// so apiMux sees /v1/ping instead of /api/v1/ping
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	// Register other top-level routes on mainMux
	mainMux.HandleFunc("/render", templateRender)
	mainMux.HandleFunc("/method-inspector", MethodInspector)
	mainMux.HandleFunc("/echo", Echo)
	mainMux.HandleFunc("/headers", HeaderDetective)
	mainMux.HandleFunc("/form", FormDecoder)
	mainMux.HandleFunc("/status", StatusCode)

	log.Fatal(http.ListenAndServe(":8080", mainMux))
	//http.NewServeMux()
}
