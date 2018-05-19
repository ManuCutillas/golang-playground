package main

import (
	"net/http"
	"io"
)

type hot int

func (r hot) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("custom-header", "this is a custom header")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/":
		io.WriteString(w, "<h1>Home page</h1>")
	case "/first":
		io.WriteString(w, "<h1>First page</h1>")
	case "/second":
		io.WriteString(w, "<h1>Second page</h1>")
	}
}

func main() {
	var p hot
	http.ListenAndServe(":8080", p)
}