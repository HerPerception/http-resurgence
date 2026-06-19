package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "Stranger"
	}
	fmt.Fprintf(w, "Greetings, %s!\n", name)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
