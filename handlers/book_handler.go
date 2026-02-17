package handlers

import (
	"bufio"
	"fmt"
	"main/service"
	"strings"

)


func SearchBooksHandler(reader *bufio.Reader) {
		fmt.Print("Enter your search query: ")
		searchQuery, _ := reader.ReadString('\n')
		searchQuery = strings.TrimSpace(searchQuery)
		fmt.Printf("Searching for books with query: %s\n", searchQuery)
		books, _ := service.SearchBooks(searchQuery)

		for _, group := range books.BookSummary {
			for _, book := range group {
				fmt.Printf("[%d] %s (Rating: %.2f)\n", book.ID, book.Title, book.Rating.Average)
			}
		}
}

func SearchBook(reader *bufio.Reader) {
		fmt.Print("Enter book ID: ")
		searchQuery, _ := reader.ReadString('\n')
		searchQuery = strings.TrimSpace(searchQuery)
		book, _ := service.SearchBook(searchQuery)
		fmt.Printf("[%d] %s (Rating: %.2f)\n", book.ID, book.Title, book.Rating.Average)
}