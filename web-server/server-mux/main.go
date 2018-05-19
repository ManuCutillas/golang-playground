package main

import (
	"io"
	"net/http"
)

func a(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>Home page</h1>")
}
func b(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>First page</h1>")
}
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>Second page</h1>")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/first", b)
	http.HandleFunc("/second", c)

	http.ListenAndServe(":8080", nil)
}

