package main

import "fmt"

type UserOld struct {
	name string
	age  int
}

type Book struct {
	title  string
	pages  int
	author *UserOld
}

func main() {
	//Указатели 1
	a := 6
	b := 10
	swap(&a, &b)
	fmt.Println(a, b)

	//Указатели 2

	p := 10
	increment(&p)
	fmt.Println(p)

	//Структуры 1
	u := UserOld{"Oleg", 13}
	fmt.Println(u)

	//Структуры 2
	setAge(&u, 15)
	fmt.Println(u)

	//Структуры 3
	book := Book{
		title: "hello world",
		pages: 100,
		//author: &UserOld{"Akakii", 30},
	}
	fmt.Println(book)
	//	fmt.Println(*book.author)

	printBookInfo(book)
}

// Указатели 1
func swap(a, b *int) {
	d := *a
	*a = *b
	*b = d

}

// Указатели 2
func increment(p *int) {
	*p++
}

func setAge(u *UserOld, newAge int) {
	u.age = newAge
}

func printBookInfo(book Book) {
	bookInfo := fmt.Sprintf("Книга: %s , стр: %d", book.title, book.pages)
	if book.author != nil {
		bookInfo += fmt.Sprintf(" автор: %s, возраст: %d", book.author.name, book.author.age)
	}
	fmt.Println(bookInfo)

}
