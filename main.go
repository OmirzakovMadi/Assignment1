package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OmirzakovMadi/Assignment1/Library"
)

func main() {
	lib := Library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите команду:")
		fmt.Println("1. Add - Добавить книгу")
		fmt.Println("2. Borrow - Заимствовать книгу")
		fmt.Println("3. Return - Вернуть книгу")
		fmt.Println("4. List - Показать доступные книги")
		fmt.Println("5. Exit - Выйти из программы")

		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "Add":
			var id, title, author string
			fmt.Print("Введите ID: ")
			fmt.Scanln(&id)
			fmt.Print("Введите название: ")
			fmt.Scanln(&title)
			fmt.Print("Введите автора: ")
			fmt.Scanln(&author)

			lib.AddBook(Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case "Borrow":
			var id string
			fmt.Print("Введите ID книги для заимствования: ")
			fmt.Scanln(&id)
			lib.BorrowBook(id)
		case "Return":
			var id string
			fmt.Print("Введите ID книги для возврата: ")
			fmt.Scanln(&id)
			lib.ReturnBook(id)
		case "List":
			fmt.Println("Доступные книги:")
			lib.ListBooks()
		case "Exit":
			fmt.Println("Завершение программы...")
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте ещё раз.")
		}
	}
}
