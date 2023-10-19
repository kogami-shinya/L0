package main

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
import "fmt"

func SetBit(n int64, pos uint, bitValue int) int64 {
	mask := int64(1 << pos) // Создаем маску, устанавливая нужный бит в 1
	clearMask := ^mask      // Создаем маску, чтобы установить нужный бит в 0

	if bitValue == 1 {
		n |= mask // Устанавливаем i-й бит в 1
	} else if bitValue == 0 {
		n &= clearMask // Устанавливаем i-й бит в 0
	}

	return n
}

func main() {
	var num int64 = 1024
	bitPos := uint(3) // Устанавливаем 4-й бит (индексация с 0)

	num = SetBit(num, bitPos, 1)
	fmt.Println(num)

	num = SetBit(num, bitPos, 0)
	fmt.Println(num)
}
