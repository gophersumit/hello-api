package main

import (
	"log"
	"net/http"

	"github.com/gophersumit/hello-api/handlers/rest"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
