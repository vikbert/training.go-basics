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
	books map[string]book
}

func createBookMap(books []book) *bookMap {
	bMap := make(map[string]book)
	for _, book := range books {
		bMap[book.isbn] = book
	}

	return &bookMap{
		books: bMap,
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
	fmt.Println("Book by isbn 193-753: ", myMap.indexBooksByIsbn("193-753"))
	fmt.Println("\nGroup via normal func:\n")
	fmt.Println("Book grouped by author: \n", myMap.groupBooksByAuthor())
	fmt.Println("Book grouped by publisher: \n", myMap.groupBooksByPublisher())

	fmt.Println("\nGroup via Closure func:\n")
	fmt.Println("Book grouped by publisher: \n", myMap.groupBooksWithClosure(authorKeyMatcher()))
	fmt.Println("Book grouped by author: \n", myMap.groupBooksWithClosure(publisherKeyMatcher()))
}

func (m *bookMap) indexBooksByIsbn(isbn string) book {
	return m.books[isbn]
}

func (m *bookMap) groupBooksByAuthor() map[string][]book {
	groupByAuthor := make(map[string][]book)
	for _, book := range m.books {
		groupByAuthor[book.author] = append(groupByAuthor[book.author], book)
	}

	return groupByAuthor
}

func (m *bookMap) groupBooksByPublisher() map[string][]book {
	groupByAuthor := make(map[string][]book)
	for _, book := range m.books {
		groupByAuthor[book.publisher] = append(groupByAuthor[book.publisher], book)
	}

	return groupByAuthor
}

func (m *bookMap) groupBooksWithClosure(keyMapper groupByFunc) map[string][]book {
	groupByAuthor := make(map[string][]book)
	for _, book := range m.books {
		matchedKey := keyMapper(book)
		groupByAuthor[matchedKey] = append(groupByAuthor[matchedKey], book)
	}

	return groupByAuthor
}

type groupByFunc func(b book) string

func authorKeyMatcher() groupByFunc {
	return func(b book) string { return b.author }
}

func publisherKeyMatcher() groupByFunc {
	return func(b book) string { return b.publisher }
}
