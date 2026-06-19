package main

import (
	"fmt"
	"net/http"
)

func MethodInspector(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintf(w, "You made a %s request.\n", method)
}
