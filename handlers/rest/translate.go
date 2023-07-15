package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gophersumit/hello-api/translation"
)

const defaultLanguage = "english"

type Resp struct {
	Language    string `json:"language,omitempty"`
	Translation string `json:"translation,omitempty"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")

	translation := translation.Translate(word, language)

	if translation == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("failed to encode response")
	}
}
