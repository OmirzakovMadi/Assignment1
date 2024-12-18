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

// NewLibrary создаёт новый экземпляр Library
func NewLibrary() *Library {
	return &Library{Collection: make(map[string]Book)}
}

// AddBook добавляет новую книгу
func (l *Library) AddBook(book Book) {
	l.Collection[book.ID] = book
	fmt.Println("Книга добавлена:", book.Title)
}

// BorrowBook заимствует книгу по ID
func (l *Library) BorrowBook(id string) {
	if book, exists := l.Collection[id]; exists && !book.IsBorrowed {
		book.IsBorrowed = true
		l.Collection[id] = book
		fmt.Println("Книга заимствована:", book.Title)
	} else {
		fmt.Println("Книга не найдена или уже заимствована.")
	}
}

// ReturnBook возвращает книгу по ID
func (l *Library) ReturnBook(id string) {
	if book, exists := l.Collection[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.Collection[id] = book
		fmt.Println("Книга возвращена:", book.Title)
	} else {
		fmt.Println("Книга не найдена или уже доступна.")
	}
}

// ListBooks выводит список доступных книг
func (l *Library) ListBooks() {
	for _, book := range l.Collection {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Название: %s, Автор: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
