package animal

import "fmt"

type Animal interface {
	sound()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d *Dog) sound() {
	fmt.Println(d.Name, "GUAU!!")

}

func (c *Cat) sound() {
	fmt.Println(c.Name, "MIAU!!")
}

func MakeSound(animal Animal) {
	animal.sound()
}
