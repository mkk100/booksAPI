package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/requests"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// APIKey is accessible to any function in this package
var APIKey string

func init() {
    // Load .env once when the package is initialized
    err := godotenv.Load()
    if err != nil {
        // If .env is missing in production (where you use real env vars), 
        // you might just log a warning instead of a fatal error.
        log.Println("No .env file found, using system environment variables")
    }

    APIKey = os.Getenv("booksAPI")
    
    if APIKey == "" {
        log.Println("Warning: booksAPI environment variable is not set")
    }
}

func queryFormatter(query string) (string){
	words := strings.Fields(query)
	return strings.Join(words, "+")
}

func SearchBooks(query string) (requests.Books, error){
	apiURL := "https://api.bigbookapi.com/search-books?api-key=" + APIKey + "&query=" + queryFormatter(query)
	resp, err := http.Get(apiURL)
	if err != nil {
		return requests.Books{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	return requests.Books{}, fmt.Errorf("failed to read body: %w", err)
}

	var books requests.Books
	if err := json.Unmarshal(body, &books); err != nil {
		return requests.Books{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return books, nil
}

func SearchBook(id string) (requests.Book, error) {
	apiURL := "https://api.bigbookapi.com/"+ id + "?api-key=" + APIKey 
	resp, err := http.Get(apiURL)
	if err != nil {
		return requests.Book{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	return requests.Book{}, fmt.Errorf("failed to read body: %w", err)
}
	var book requests.Book
	if err := json.Unmarshal(body, &book); err != nil {
		return requests.Book{}, fmt.Errorf("failed to parse response: %w", err)
	}
	return book, nil
}
