package main

import (
	"fmt"
)

type book struct {
	name      string
	author    string
	pageCount int
	price     int
}

//class lardagi constructor o'rniga ishlatilar ekan
func newBook(name, author string, pageCount, price int) book {
	return book{
		name:      name,
		author:    author,
		pageCount: pageCount,
		price:     price,
	}
}

//method larni shu ko'rinishda yozish mumkin ekan
func (e *book) getInfo() string {
	return fmt.Sprintf("name: %s\npage count: %d\nprice: %d", e.name, e.pageCount, e.price)
}

//new method though we can update book's price
func (e *book) setPrice(price int) {
	e.price = price
}

func main() {
	book1 := newBook("Algorithms and Data Structures in GO", "Maksim", 104, 10)

	fmt.Printf("%+v\n\n", books)

	book1.setPrice(50)
	fmt.Println(book1.getInfo())
}

var books = []book{
	{
		name:      "Algorithms and Data Structures in GO",
		author:    "Zhashkevych",
		pageCount: 104,
		price:     10,
	},
	{
		name:      "Golang Web Development",
		author:    "Anvar Shomurodov",
		pageCount: 210,
		price:     20,
	},
}
