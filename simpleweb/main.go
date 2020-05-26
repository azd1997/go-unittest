package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/mail", handleMail)
	mux.HandleFunc("/query", handleQuery)

	http.ListenAndServe(":8000", mux)
}
