package main

//Реализовать собственную функцию sleep.
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начало выполнения программы")

	mySleep(3) // Вызываем функцию mySleep с ожиданием в 3 секунды

	fmt.Println("Программа продолжает выполнение после ожидания")
}

// Функция mySleep имитирует ожидание заданное количество секунд
func mySleep(seconds int) {
	// Текущее время
	currentTime := time.Now()

	// Добавляем заданное количество секунд к текущему времени
	targetTime := currentTime.Add(time.Duration(seconds) * time.Second)

	// Цикл, пока текущее время меньше целевого времени
	for currentTime.Before(targetTime) {
		currentTime = time.Now()
	}
}
