package controller

import (
	"encoding/json"
	"net/http"
)

// LongURLHandler handles URL shortening.
// @Summary Shorten a URL
// @Tags URL Shortener
// @Accept json
// @Produce json
// @Param req body Request true "URL to be long"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /long-url [post]

func (s *ShortenURLController) LongURLHandler(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		ShortURL string `json:"short_url"`
	}

	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	originUrl, found := s.urlMap[req.ShortURL]
	if !found {
		http.Error(w, "short URL is not found", http.StatusNotFound)
		return
	}

	response := map[string]string{
		"url":       originUrl,
		"short_url": req.ShortURL,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
