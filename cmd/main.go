package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		resp := Resp{
			Language:    "English",
			Translation: "Hello World!",
		}

		if err := enc.Encode(resp); err != nil {
			panic("failed to encode response")

		}
	})

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language,omitempty"`
	Translation string `json:"translation,omitempty"`
}
