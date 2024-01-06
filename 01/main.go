// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import "fmt"

// Родительская структура
type Human struct {
	Name	string
	Surname string
}

func (h *Human) Info() {
	fmt.Printf("I am %s %s\n", h.Name, h.Surname)
}

// Дочерняя структура
type Action struct {
	Human // встраивание структуры, аналог наследования
	Exercise string
}

func (a *Action) Do() {
	fmt.Printf("%s %s is %s now\n", a.Name, a.Surname, a.Exercise)
}



func main() {
	a := Action{
		Human: Human{
			Name: "Test",
			Surname: "Testov",
		},
		Exercise: "jumping",
	}

	a.Info()
	a.Do()
}