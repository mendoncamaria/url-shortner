package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func URLRoutes(r *chi.Router, db *sql.DB) {
	r.Route("/urls", func(r *chi.Router) {
		r.Post("/", createURL(db))
		r.Get("/{shortCode}", getURL(db))
	})
}

// ... implementation of createURL and getURL functions
