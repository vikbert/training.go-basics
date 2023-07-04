package main

import (
	"fmt"
)

type book struct {
	isbn      string
	author    string
	publisher string
}

type bookMap struct {
	books []book
}

func createBookMap(books []book) *bookMap {
	return &bookMap{
		books: books,
	}
}

func main() {
	fmt.Println("Maps:")
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}

	myMap := createBookMap(books[:])
	fmt.Println("Book by isbn 193-753: ", myMap.getBookByIsbn("193-753"))

	fmt.Println("\nGroup via Closure func:\n")
	resolveAuthor := func(b book) string { return b.author }
	resolvePublisher := func(b book) string { return b.publisher }
	fmt.Println("Book grouped by publisher: \n", myMap.groupBookByDefinition(resolveAuthor))
	fmt.Println("Book grouped by author: \n", myMap.groupBookByDefinition(resolvePublisher))
}

func (m *bookMap) getBookByIsbn(isbn string) book {
	bMap := make(map[string]book)
	for _, book := range m.books {
		bMap[book.isbn] = book
	}

	return bMap[isbn]
}

func (m *bookMap) groupBookByDefinition(fn func(book) string) map[string][]book {
	booksMap := make(map[string][]book)
	for _, book := range m.books {
		groupKey := fn(book)
		booksMap[groupKey] = append(booksMap[groupKey], book)
	}

	return booksMap
}
