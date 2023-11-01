package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte("This is main page."))
	})

	http.ListenAndServe(":8800", nil)
}
