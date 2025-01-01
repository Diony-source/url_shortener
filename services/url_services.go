package services

import (
	"errors"
	"url_shortener/database"
	"url_shortener/utils"
)

func CreateShortURL(longURL string) (string, error) {
	shortURL := utils.GenerateShortURL(6)
	err := database.SaveURL(shortURL, longURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func GetLongURL(shortURL string) (string, error) {
	longURL, err := database.FetchURL(shortURL)
	if err != nil {
		return "", errors.New("short URL not found")
	}
	return longURL, nil
}
