package main

import (
	"bufio"
	"fmt"
	"main/handlers"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Bookstore")

	for {
		fmt.Println("1. Search Books (search)")
		fmt.Println("2. Add to the cart (add {bookId})")
		fmt.Println("3. Delete book (delete {bookId})")
		fmt.Println("4. Search By Id (searchById {bookId})")
		fmt.Println("5. Quit the app (quit)")
		fmt.Print("\nYour action: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
		}

		query := strings.TrimSpace(strings.ToLower(input))

		if query == "quit" {
			break
		}		
		if query == "searchbyid" {
			handlers.SearchBook(reader)
		}
		if query == "search" {
			handlers.SearchBooksHandler(reader)
		}

	}
}