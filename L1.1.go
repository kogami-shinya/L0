package main

//Условие: Дана структура Human (с произвольным набором полей и методов).Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
//Так как в Golang отсутствует наследование, нельзя напрямую встраивать методы из одной структуры в другую. Однако, можно достичь подобного поведения, объединив две структуры в композицию.
import "fmt"

type Human struct {
	FirstName string
	LastName  string
}

func (h Human) Speak() {
	fmt.Println("I am a human")
}

type Action struct {
	Human
}

func (a Action) Run() {
	fmt.Println("Running")
}

func main() {
	h := Action{
		Human: Human{
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	/*
		h.Human.FirstName = "John"
		h.Human.LastName = "Doe"

		h.FirstName = "John"
		h.LastName = "Doe"
	*/

	fmt.Println(h.FirstName, " ", h.LastName)
	h.Speak() // вызов метода Speak из структуры Human
	h.Run()   // вызов метода Run из структуры Action
}
