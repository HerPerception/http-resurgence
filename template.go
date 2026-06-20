// EXERCISE 7


package main

import (
	"net/http"
	"text/template"
)

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  <h1>{{.Title}}</h1>
{{if eq .Style "bold"}}<strong>{{.Body}}</strong>{{else}}{{.Body}}{{end}}
</body>
</html>
`

type PageData struct {
	Title string
	Body  string
	Style string
}

var tmpl = template.Must(template.New("page").Parse(tmplStr))

func templateRender(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	style := r.URL.Query().Get("style")
	if len(title) == 0 || len(body) == 0 {
		http.Error(w, "title and body are required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err := tmpl.Execute(w, PageData{Title: title, Body: body, Style: style})
	if err != nil {
		http.Error(w, "template execution failed", http.StatusInternalServerError)
		return
	}
}

/*
When does template.Must panics?
template.Must panics when an error occurs while parsing file(s).
*/
