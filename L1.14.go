package main

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
import (
	"fmt"
	"reflect"
)

func getType(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("Тип: int")
	case reflect.String:
		fmt.Println("Тип: string")
	case reflect.Bool:
		fmt.Println("Тип: bool")
	case reflect.Chan:
		fmt.Println("Тип: channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var a int = 10
	var b string = "Hello"
	var c bool = true
	var d = make(chan int)

	getType(a) // Вывод: Тип: int
	getType(b) // Вывод: Тип: string
	getType(c) // Вывод: Тип: bool
	getType(d) // Вывод: Тип: channel
}
