package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortenURL   string    `json:"shorten_url"`
	CreationDate time.Time `json:"creation_time"`
}

var urlDB = make(map[string]URL)

func main() {
	fmt.Println("Starting URL-Shortener...")

	http.HandleFunc("/", rootPage)
	http.HandleFunc("/short", shortURLhandler)
	http.HandleFunc("/redirect/", redirectURLhandler)

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error while starting server")
		panic(err)
	}
}

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("Final shortened:", hash[:5])
	return hash[:5]
}

func createURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL

	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortenURL:   shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func shortURLhandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)

	response := struct {
		Shortened string `json:"short_url"`
	}{Shortened: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLhandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
