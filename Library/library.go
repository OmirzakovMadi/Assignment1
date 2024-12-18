package Library

import "fmt"

type Book struct {
    ID         string
    Title      string
    Author     string
    IsBorrowed bool
}

type Library struct {
    Collection map[string]Book
}

func NewLibrary() *Library {
    return &Library{Collection: make(map[string]Book)}
}

// AddBook добавляет новую книгу в коллекцию.
// Если книга с таким ID уже существует, она будет перезаписана.
func (l *Library) AddBook(book Book) {
    l.Collection[book.ID] = book
}

// BorrowBook устанавливает IsBorrowed = true для книги с данным ID.
// Если книги нет или она уже заимствована, выведем сообщение об ошибке.
func (l *Library) BorrowBook(id string) {
    book, exists := l.Collection[id]
    if !exists {
        fmt.Println("Книга с таким ID не найдена!")
        return
    }
    if book.IsBorrowed {
        fmt.Println("Книга уже заимствована.")
        return
    }
    book.IsBorrowed = true
    l.Collection[id] = book
    fmt.Println("Книга заимствована.")
}

// ReturnBook устанавливает IsBorrowed = false для книги с данным ID.
// Если книги нет или она не была заимствована, выведем сообщение.
func (l *Library) ReturnBook(id string) {
    book, exists := l.Collection[id]
    if !exists {
        fmt.Println("Книга с таким ID не найдена!")
        return
    }
    if !book.IsBorrowed {
        fmt.Println("Книга и так была доступна.")
        return
    }
    book.IsBorrowed = false
    l.Collection[id] = book
    fmt.Println("Книга возвращена.")
}

// ListBooks выводит список всех книг.
// Если хотите отфильтровать только доступные книги, можно проверять IsBorrowed.
func (l *Library) ListBooks() {
    if len(l.Collection) == 0 {
        fmt.Println("Нет книг в библиотеке.")
        return
    }
    for _, book := range l.Collection {
        status := "Доступна"
        if book.IsBorrowed {
            status = "Заимствована"
        }
        fmt.Printf("ID: %s, Название: %s, Автор: %s, Статус: %s\n",
            book.ID, book.Title, book.Author, status)
    }
}
