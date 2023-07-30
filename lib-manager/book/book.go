package book

import "fmt"

type Printable interface {
	PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// NewBook Simula constructor
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nlevel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}

func Print(p Printable) {
	p.PrintInfo()
}
