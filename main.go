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
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\nВыберите действие:")
        fmt.Println("1. Add (Добавить книгу)")
        fmt.Println("2. Borrow (Заимствовать книгу)")
        fmt.Println("3. Return (Вернуть книгу)")
        fmt.Println("4. List (Список всех книг)")
        fmt.Println("5. Exit (Выйти)")
        fmt.Print("Введите номер операции: ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "1":
            // Добавление книги
            fmt.Print("Введите ID книги: ")
            id, _ := reader.ReadString('\n')
            id = strings.TrimSpace(id)

            fmt.Print("Введите название книги: ")
            title, _ := reader.ReadString('\n')
            title = strings.TrimSpace(title)

            fmt.Print("Введите автора книги: ")
            author, _ := reader.ReadString('\n')
            author = strings.TrimSpace(author)

            // По умолчанию IsBorrowed = false при добавлении
            book := Library.Book{
                ID:         id,
                Title:      title,
                Author:     author,
                IsBorrowed: false,
            }
            lib.AddBook(book)
            fmt.Println("Книга добавлена.")

        case "2":
            // Заимствование книги
            fmt.Print("Введите ID книги для заимствования: ")
            id, _ := reader.ReadString('\n')
            id = strings.TrimSpace(id)
            lib.BorrowBook(id)

        case "3":
            // Возврат книги
            fmt.Print("Введите ID книги для возврата: ")
            id, _ := reader.ReadString('\n')
            id = strings.TrimSpace(id)
            lib.ReturnBook(id)

        case "4":
            // Список всех книг
            lib.ListBooks()

        case "5":
            // Выход
            fmt.Println("Программа завершена.")
            return

        default:
            fmt.Println("Неверный выбор. Попробуйте еще раз.")
        }
    }
}
