package main

import (
    "database/sql"
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/mattn/go-sqlite3"
    "url-shortener/api/models"
    "url-shortener/api/routes"
)

func main() {
    db, err := sql.Open("sqlite3", "urls.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create a table if it doesn't exist
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY AUTOINCREMENT, original_url TEXT, short_code TEXT UNIQUE)")
    if err != nil {
        panic(err)
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    routes.URLRoutes(r, db)

    http.ListenAndServe(":8080", r)
}