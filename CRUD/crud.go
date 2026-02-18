package CRUD

import (
	"fmt"
	"main/requests"
	"main/service"
	"strconv"
)

var library = map[string]requests.Book{}

func AddBook(id string) error {
	book, err := service.SearchBook(id)
	if err != nil {
		return err
	}
	library[strconv.Itoa(book.ID)] = book
	return nil
}

func DeleteBook(id string) {
	delete(library, id)
}

func ViewBooks() {
	fmt.Println(len(library))
	for _, book := range library {
		fmt.Printf("[%d] %s (Rating: %.2f)\n", book.ID, book.Title, book.Rating.Average)
	}
}