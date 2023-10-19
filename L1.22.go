package main

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2000000)
	b := big.NewInt(2000000)

	// Перемножение
	product := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", product)

	// Деление с округлением
	quotient := new(big.Int).Quo(a, b)
	fmt.Println("Частное:", quotient)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum)

	// Вычитание
	difference := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", difference)
}
