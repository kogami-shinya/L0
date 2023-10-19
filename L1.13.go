package main

//Поменять местами два числа без создания временной переменной.
import "fmt"

func swapWithoutTemp(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}

func main() {
	x, y := 10, 20
	x, y = swapWithoutTemp(x, y)
	fmt.Println(x, y)
}
