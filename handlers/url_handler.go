package handlers

import (
	"net/http"
	"url_shortener/entities"
	"url_shortener/services"
	"url_shortener/utils"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url entities.URL
	if err := utils.DecodeJSON(r, &url); err != nil {
		utils.WriteErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := services.CreateShortURL(url.LongURL)
	if err != nil {
		utils.WriteErrorResponse(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(w, map[string]string{"short_url": shortURL}, http.StatusOK)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/redirect/"):]

	longURL, err := services.GetLongURL(shortURL)
	if err != nil {
		utils.WriteErrorResponse(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
