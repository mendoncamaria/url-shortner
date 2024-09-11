package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"url-shortner/api/api/models" //right import, don't change this

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func URLRoutes(r *chi.Router, db *sql.DB) {
	r.Post("/", createURL(db))
	r.Get("/{shortCode}", getURL(db))
}

func createURL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			OriginalURL string `json:"original_url"`
		}

		// Decode the request body using encoding/json
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			render.Render(w, r, render.BadRequest(err))
			return
		}

		shortCode := models.generateShortCode()
		url := models.URL{
			OriginalURL: data.OriginalURL,
			ShortCode:   shortCode,
		}

		err := url.Save(db)
		if err != nil {
			render.Render(w, r, render.ServerError(err))
			return
		}

		render.JSON(w, r, url)
	}
}

func getURL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := chi.URLParam(r, "shortCode")

		originalURL, err := models.GetOriginalURLByShortCode(db, shortCode)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) { // Handle case of non-existent URL
				render.Render(w, r, render.NotFound(err))
			} else {
				render.Render(w, r, render.ServerError(err))
			}
			return
		}

		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}
