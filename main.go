package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/buy", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "OK")
	})

	http.ListenAndServe(":3000", nil)
}
