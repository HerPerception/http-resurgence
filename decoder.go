package main

import (
	"fmt"
	"net/http"
)

func FormDecoder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	lang := r.FormValue("language")
	if len(username) == 0 {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	} else if len(lang) == 0 {
		http.Error(w, "language is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s, you are coding in %s!\n", username, lang)
}
