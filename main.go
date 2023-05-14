package main

import (
	"net/http"
)

func main() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "version: 0.1.2\nYour Request Path is " + req.URL.Path
	w.Write([]byte(str))
}
