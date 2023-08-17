package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ShortenURLHandler handles URL shortening.
// @Summary Shorten a URL
// @Tags URL Shortener
// @Accept json
// @Produce json
// @Param req body Request true "URL to be shortened"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /short-url [post]

func (s *ShortenURLController) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		URL string `json:"url"`
	}

	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	shortURL := fmt.Sprintf("http://shorten-url-service/short/%s", generatedShortURL())
	s.urlMap[shortURL] = req.URL

	response := map[string]string{
		"url":       req.URL,
		"short_url": shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generatedShortURL() string {
	return "aBCsa"
}
