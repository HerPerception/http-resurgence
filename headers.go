/*r.Header.Get returns an empty string for a header key that was never sent.*/
package main

import (
	"fmt"
	"net/http"
)

func HeaderDetective(w http.ResponseWriter, r *http.Request) {
	/*r.Header.Get() is case-insensitive so the key "X-Custom-Token" is read same as "x-custom-token".*/
	token := r.Header.Get("X-Custom-Token")
	content := r.Header.Get("Content-Type")
	if len(content) == 0 {
		content = "Content-Type not provided"
	}
	if len(token) == 0 {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Token: %s\n Content-Type: %s\n", token, content)
}
