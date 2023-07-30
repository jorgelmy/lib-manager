package main

import (
	"fmt"
	"lib-manager/animal"
	"lib-manager/book"
)

func main() {
	myBook := book.NewBook("Harry Potter", "Fulana D.", 984)

	myBook.PrintInfo()

	myBook.SetTitle("Harry Potter: El prisionero de Azkaban")
	fmt.Println(myBook.GetTitle())

	myTextbook := book.NewTextBook("HP: Piedra filosofal", "Fulana D. A.", 564, "Editorialisima", "5")
	myTextbook.PrintInfo()

	fmt.Println("_______________________________________")
	book.Print(myBook)
	book.Print(myTextbook)

	fmt.Println("_______________________________________")
	gato := animal.Cat{"Micho"}
	perro := animal.Dog{Name: "Firu"}

	animal.MakeSound(&gato)
	animal.MakeSound(&perro)

}
