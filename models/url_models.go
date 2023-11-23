package models

import (
	"go-url-shortener/database"
	"time"
)

type UrlRequestBody struct {
	LongUrl  string `json:"url"`
	ShortUrl string `json:"short_url"`
}

type Url struct {
	LongUrl      string    `json:"long_url" db:"long_url"`
	ShortUrl     string    `json:"short_url" db:"short_url"`
	AccessCount  int       `json:"access_count" db:"access_count"`
	LastAccessed time.Time `json:"last_accessed" db:"last_accessed"`
	CreateAt     time.Time `json:"create_at" db:"create_at"`
	UpdateAt     time.Time `json:"update_at" db:"update_at"`
}

func CreateShortenUrl(longUrl string, shortUrl string) error {
	query := `INSERT INTO url (long_url, short_url, access_count) VALUES ($1, $2, $3)`

	_, err := database.DBClient.Exec(query, longUrl, shortUrl, 0)
	if err != nil {
		return err
	}

	return nil
}

func GetURLByShortURL(shortUrl string) (Url, error) {
	var url Url

	query := `SELECT long_url, short_url, access_count, 
	last_accessed, create_at, update_at FROM url WHERE short_url = $1`

	err := database.DBClient.Get(&url, query, shortUrl)
	if err != nil {
		return url, err
	}

	return url, nil
}

func UpdateAccessCount(shortUrl string) error {
	query := `UPDATE url SET access_count = access_count + 1, last_accessed = $1 WHERE short_url = $2`

	_, err := database.DBClient.Exec(query, time.Now(), shortUrl)
	if err != nil {
		return err
	}

	return nil
}
