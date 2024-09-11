package models

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"
)

const shortCodeLength = 6

type URL struct {
	ID          int
	OriginalURL string
	ShortCode   string
}

func (u *URL) Save(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO urls (original_url, short_code) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.OriginalURL, u.ShortCode)
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalURLByShortCode(db *sql.DB, shortCode string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_code = ?", shortCode).Scan(&originalURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("URL not found")
		}
		return "", err
	}
	return originalURL, nil
}

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, shortCodeLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
