package main

import (
	"errors"
	"fmt"
	"local/042_methods-interfaces/a_methods/stack"
)

type book struct {
	isbn      string
	author    string
	publisher string
}

type Bookshelf struct {
	books []book
}

func main() {
	fmt.Println("Methods:")
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}

	bs := Bookshelf{books[:]}
	bs.add(book{
		isbn:      "999-999",
		author:    "T. Hanks",
		publisher: "Springer",
	})

	fmt.Println(bs)
	fmt.Println(bs.forIsbn("999-999"))
	fmt.Println(bs.all())

	// stack demo
	myStack := stack.NewStack()
	myStack.Push(0)
	myStack.Push(1)
	myStack.Push(2)

	fmt.Println("Stack:", myStack)
	fmt.Println("Stack::Pop:", myStack.Pop())
	fmt.Println("Stack::size:", myStack.Size())
	fmt.Println("Stack:", myStack)
	fmt.Println("Stack::Peek:", myStack.Peek())
	fmt.Println("Stack::Size:", myStack.Size())
	fmt.Println("Stack:", myStack)
	fmt.Println("Stack::Size:", myStack.Size())
}

func (bs *Bookshelf) add(b book) {
	bs.books = append(bs.books, b)
}

func (bs *Bookshelf) forIsbn(isbn string) (book, error) {
	var found book
	for _, b := range bs.books {
		if b.isbn == isbn {
			return b, nil
		}
	}

	return found, errors.New("not found")
}

func (bs *Bookshelf) all() []book {
	return bs.books
}
