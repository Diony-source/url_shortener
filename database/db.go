package database

import (
	"errors"
	"sync"
)

var urlStore = make(map[string]string)
var mu sync.RWMutex

func SaveURL(shortURL, longURL string) error {
	mu.Lock()
	defer mu.Unlock()
	urlStore[shortURL] = longURL
	return nil
}

func FetchURL(shortURL string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()
	longURL, exists := urlStore[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}
	return longURL, nil
}
