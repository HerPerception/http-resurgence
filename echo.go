// EXERCISE 2
package main

import (
	"io"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		defer r.Body.Close()
		return
	}
	/*When I did my check as 'if len(data) == 0', POST requests with the body empty were rejected,
	  but when I used 'if data == nil', the requests went through and content-length was returned at zero.
	  I eventually opted for len(data) == 0 because the task explicitly stated that as the instruction.*/
	if len(data) == 0 {
		http.Error(w, "body cannot be empty", http.StatusMethodNotAllowed)
		return
	}
	//fmt.Fprintln(w, string(data))
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
	/*
		When I called w.Header().Set() before w.Write(), This was the output with curl:

		http-resurgence git:(main) ✗ curl -i -X POST http://localhost:8080/echo -d "Hello"
		HTTP/1.1 200 OK
		Content-Type: text/plain
		Date: Fri, 19 Jun 2026 14:15:48 GMT
		Content-Length: 5

		Hello%

		When I called w.Write() before w.Header().Set(), this was the output with curl:

		http-resurgence git:(main) ✗ curl -i -X POST http://localhost:8080/echo -d "Hello"
		HTTP/1.1 200 OK
		Date: Fri, 19 Jun 2026 14:19:23 GMT
		Content-Length: 5
		Content-Type: text/plain; charset=utf-8

		Hello%

		Do you see the difference?
		The content-type  came before at the first instance, then it came last at the second instance. Interesting, right? I just learned it today too.

		What happens to r.Body if you read it twice without closing it?
	*/
}
