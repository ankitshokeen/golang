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
	OrignalURL   string    `json:"orignal_url"`
	ShortenURL   string    `json:"shorten_url"`
	CreationDate time.Time `json:"creation_time"`
}

var urlDB = make(map[string]URL)

func main() {
	fmt.Println("Starting URL-Shortner...")

	// var urlToBeConverted = "https://github.com/ankitshokeen/golang"
	// generateShortURL(urlToBeConverted)

	http.HandleFunc("/", rootPage)
	http.HandleFunc("/short", shortURLhandler)

	fmt.Println("starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error while starting server")
		panic(err)
	}

}

func generateShortURL(OrignalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OrignalURL))
	fmt.Println("hashed value: ", hasher)

	data := hasher.Sum(nil)
	fmt.Println("hashed data in slice : ", data)

	hash := hex.EncodeToString(data)
	fmt.Println("encoded to string: ", hash)

	fmt.Println("final shortened: ", hash[:5])
	return hash[:5]
}

func createURL(OrignalURL string) string {
	shortURL := generateShortURL(OrignalURL)
	id := shortURL

	urlDB[id] = URL{
		ID:           id,
		OrignalURL:   OrignalURL,
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

	shortURL := generateShortURL(data.URL)
	// fmt.Fprintf(w, shortURL)

	response := struct {
		Shortend string `json:"short_url"`
	}{Shortend: shortURL}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLhandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
	}

	http.Redirect(w, r, url.OrignalURL, http.StatusNotFound)
}
