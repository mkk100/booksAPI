package main

import (
	"bufio"
	"fmt"
	"main/CRUD"
	"main/handlers"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Bookstore")

	for {
		fmt.Print("\n")
		fmt.Println("1. Search Books (search)")
		fmt.Println("2. Add to the cart (add {bookId})")
		fmt.Println("3. Delete book (delete {bookId})")
		fmt.Println("4. Search By Id (searchById {bookId})")
		fmt.Println("5. View the cart (view cart)")
		fmt.Println("6. Quit the app (quit)")
		fmt.Print("\nYour action: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
		}

		query := strings.TrimSpace(strings.ToLower(input))
		parts := strings.Fields(query)
		if len(parts) == 0 {
			continue
		} else if len(parts) == 1 {
		if query == "quit" {
			break
		}		
		if query == "searchbyid" {
			handlers.SearchBookHandler(reader)
		}
		if query == "search" {
			handlers.SearchBooksHandler(reader)
		}
		} else{
			if parts[0] == "add" {
				CRUD.AddBook(parts[1])
			}
			if parts[0] == "delete" {
				CRUD.DeleteBook(parts[1])
			}
			if parts[0] == "view" {
				CRUD.ViewBooks()
			}
		}
	}
}